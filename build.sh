#!/bin/sh

export GOPATH=`pwd`

go build -o TextMiningCompiler compiler
go build -o TextMiningApp app
