package main

import (
	"fmt"
	"os"

	checkmarx "github.com/SAP/jenkins-library/pkg/checkmarx"
	format "github.com/SAP/jenkins-library/pkg/format"
)

func main() {
	fmt.Println("Piper Checkmarx XML to SARIF converter standalone")
	fmt.Println("Run: ./converter filename.xml")
	if len(os.Args[1:]) == 0 {
		fmt.Println("need filename as first argument")
		return
	}
	cxxmlFileName := os.Args[1]
	var sarif format.SARIF
	var err error
	fmt.Println("Running converter...")
	sarif, err = checkmarx.ConvertCxxmlToSarif(nil, cxxmlFileName, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = checkmarx.WriteSarif(sarif)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Report written under ./checkmarx/result.sarif")
}
