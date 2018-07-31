#!/bin/sh

OLD=$GOPATH

export GOPATH=`pwd`

go build -o TextMiningCompiler main_compiler
echo 'TextMiningCompiler generated'

go build -o TextMiningApp main_app
echo 'TextMiningApp generated'

export GOPATH=$OLD
