#!/bin/bash

# requires a working golang setup
# build tested only on linux, can target all platforms

PROG=spew
VER=$(date -d @$(git log -1 --format=%ct) '+%Y%m%d')
for i in linux darwin windows
do
    echo "building $VER of $PROG into bin/$i directory for $i..."
    mkdir -p bin/$i
    CGO_ENABLED=0 GOOS=$i GOARCH=amd64 go build -o bin/$i/spew -ldflags "-X main.v=$VER" .
done
    


