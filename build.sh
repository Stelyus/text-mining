#!/bin/sh

OLD=$GOPATH

export GOPATH=`pwd`

go build -o TextMiningCompiler main_compiler
echo 'TextMiningCompiler done'

go build -o TextMiningApp main_app
echo 'TextMiningApp done'

export GOPATH=$OLD
