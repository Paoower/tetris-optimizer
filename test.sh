#!/bin/bash

FILES=("bad0.txt" "bad1.txt" "bad2.txt" "bad3.txt" "bad4.txt" "badf.txt" "good0.txt" "good1.txt" "good2.txt" "good3.txt" "hard.txt")

for file in "${FILES[@]}"
do
    echo "Testing file: $file"
    go run . tetromino/"$file"
    echo "-------------------"
done
