#!/bin/sh

go test -bench . -run '^$' -count 5 > result.txt
go run main.go -pat 'strings.Builder,bytes.Buffer' result.txt
