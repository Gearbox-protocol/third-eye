#!/bin/zsh

for i in `ls jsonnet/**/*.jsonnet`; do
    echo $i
    jsonnetfmt --in-place $i
    dir=`dirname $i`
    mkdir -p inputs/${dir#*/}
    outputFile=`echo "inputs/${i#*/}" | cut -f 1 -d '.'`
    jsonnet $i --output-file "$outputFile.json"
done