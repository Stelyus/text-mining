#!/bin/sh

OLD=$GOPATH

export GOPATH=`pwd`

go build -o TextMiningCompiler compiler
go build -o TextMiningApp app

export GOPATH=$OLD