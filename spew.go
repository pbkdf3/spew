/**
  Copyright (C) 2017 ADP, LLC

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

  by Daniel Reznick

*/

package main

import (
	"bufio"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/jawher/mow.cli"
)

/* randstring from stack overflow, slightly modified */
const letterBytes = ` !"#$%&\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~`
const (
	letterIdxBits = 7                    // 7 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandString(src rand.Source, n int) []byte {
	b := make([]byte, n) // XXX NOT extra space for potential EOL later appended
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

/* goroutine, exits when program exits (messy but in this case ok?) */
func spew(out chan []byte) {

	src := rand.NewSource(time.Now().UnixNano())

	for {
		out <- append(RandString(src, *ll), eol...)
	}
}

var l, ll *int
var eol string
var v string // filled in by linker
func main() {
	app := cli.App("spew", "generate random strings, one per line")
	app.Spec = "[LENGTH [LINES]]"
	app.Version("v version", v)

	ll = app.IntArg("LENGTH", 32, "length of generated string")
	l = app.IntArg("LINES", 0, "number of lines of output (0 will output a string without a newline)")

	eol = "\n"
	app.Action = func() {
		if *l < 1 {
			eol = ""
			*l = 1
		}

		f := bufio.NewWriter(os.Stdout)
		defer f.Flush()
		out := make(chan []byte, 1024)

		var goroutines int
		/* run on up to 8 threads */
		if runtime.GOMAXPROCS(-1) >= 8 {
			goroutines = 8
		} else {
			goroutines = runtime.GOMAXPROCS(-1)
		}

		for i := 0; i < goroutines; i++ {
			go spew(out)
		}

		n := 0
		for rnd := range out {
			f.Write(rnd)
			n++
			if n >= *l {
				break
			}
		}

	}
	app.Run(os.Args)
}
