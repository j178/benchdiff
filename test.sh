#!/bin/sh

go test -bench . -benchmem -run '^$' -count 5 > result.txt
go run main.go -name 'Builder-vs-Buffer' -pat 'strings.Builder,bytes.Buffer' result.txt
