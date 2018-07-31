#!/bin/sh

OLD=$GOPATH

export GOPATH=`pwd`

go build -o TextMiningCompiler main_compiler
go build -o TextMiningApp app

export GOPATH=$OLD
