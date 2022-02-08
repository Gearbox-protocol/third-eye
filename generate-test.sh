#!/bin/zsh

for i in `ls jsonnet/**/*.jsonnet`; do
    echo $i
    dir=`dirname $i`
    mkdir -p tests/${dir#*/}
    outputFile=`echo "tests/${i#*/}" | cut -f 1 -d '.'`
    jsonnet $i --output-file "$outputFile.json"
done