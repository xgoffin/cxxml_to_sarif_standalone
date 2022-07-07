# CxXML to SARIF standalone

## Usage

Can be used locally, lightweight, no audit information

`go run main.go CxSASTResults_[...].xml`

`go run .\main.go CxSASTResults_2022_02_10T11_25_53_82666638Z.xml`

File can be found under `checkmarx/result.sarif`

## Compiling & running 

Seems to give faster results

`go build -o converter`

`./converter CxSASTResults_[...].xml`

## Caveat

Running the converter twice in a row will replace the output file.