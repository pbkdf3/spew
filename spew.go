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
	"math/rand"
	"os"
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

var src = rand.NewSource(time.Now().UnixNano())

func RandString(n int) string {
	b := make([]byte, n)
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

	return string(b)
}

func main() {
	app := cli.App("spew", "generate random strings, one per line")
	app.Spec = "[LENGTH [LINES]]"
	ll := app.IntArg("LENGTH", 32, "length of generated string")
	l := app.IntArg("LINES", 0, "number of lines of output (0 will output a string without a newline)")

	eol := "\n"
	app.Action = func() {
		if *l < 1 {
			eol = ""
			*l = 1
		}
		for i := 0; i < *l; i++ {
			os.Stdout.Write([]byte(RandString(*ll) + eol))
		}
	}
	app.Run(os.Args)
}
