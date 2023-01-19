package main

import (
	"fmt"

	"example.com/go-demo-1/filesInputOutput"
	"example.com/go-demo-1/parserCSVFile"
	"example.com/go-demo-1/printFirstLine"
)

func main() {
	fmt.Println("Start project")
	fmt.Println()

	x := 100
	fileName := "input.csv"

	data := filesInputOutput.ReturnData(fileName)

	printFirstLine.PrintFields()

	allPersonalDetails := parserCSVFile.Parser(data, x)
	filesInputOutput.WriteFiles(allPersonalDetails)

	fmt.Println()
	fmt.Println("End project")
}
