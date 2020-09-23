#!/bin/bash

set -e

if [ "$1" != "" ];
then
    f=$(go list ./... | grep -v 'examples' | grep $1)
else
    f=$(go list ./... | grep -v 'examples')
fi

tfile="coverage.txt"
hfile="coverage.html"
echo "" > $tfile
echo "" > $hfile

go test -coverprofile=$tfile -covermode=atomic $f
go tool cover -html=$tfile -o $hfile
