# CSV Repeater
This utility command tool will use given CSV file to generate an output with
terms being repeated by number of times given.

## Source CSV requirements
* Must have at least two columns
* Columns are delimited by comma (`","`)
* File **should not** contain any headlines/column name rows
* Values must **not be quoted**

## Example of use, MVP
```bash
you@machine:~ cat myfile.csv
keywordA,1
keywordB,2
keywordC,3
you@machine:~ csvrepeater myfile.csv
keywordB
keywordC
keywordA
keywordC
keywordC
keywordB
```

## Development ideas/improvements
* Allow specifying if randomisation is wanted (`--random`, default `"true"`)
* Allow specifying delimiter (`--delimiter`, default `","`)
* Allow specifying the column source for term (`--columnNumTerm`, default `0`)
* Allow specifying the column source for repeat number (`--columnNumRepeat`, default `1`)
* Allow skipping number of heading lines (`--skipHeading`, default `0`)
* Allow using multiple files
