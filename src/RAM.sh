#!/bin/bash
for i in `seq 1 10`; do
  gtime -v go run main.go file.go radix.go ../ressources/words.txt dict.bin
done
