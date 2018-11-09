# CSV Column Repeater
This utility command tool will use given CSV file to generate an output with
terms being repeated by number of times given.

## Source CSV requirements
* Must have at least two columns
* Columns are delimited by comma (`","`)
* File **should not** contain any headlines/column name rows
* Values must **not be quoted**

## Getting started

* Setup your go environment (for example in `~/go`)

```bash
you@machine:~ cd ~/go
you@machine:~/go go get github.com/wunderio/csvcolumnrepeater
you@machine:~/go go install github.com/wunderio/csvcolumnrepeater
you@machine:~/go cd ~
you@machine:~ cat myfile.csv
keywordA,1
keywordB,2
keywordC,3
you@machine:~ csvcolumnrepeater myfile.csv
keywordA
keywordB
keywordC
keywordB
keywordC
keywordC
```

## Development ideas/improvements
* Allow specifying if randomisation is wanted (`--random`, default `"true"`)
* Allow specifying delimiter (`--delimiter`, default `","`)
* Allow specifying the column source for term (`--columnNumTerm`, default `0`)
* Allow specifying the column source for repeat number (`--columnNumRepeat`, default `1`)
* Allow skipping number of heading lines (`--skipHeading`, default `0`)
* Allow using multiple files
