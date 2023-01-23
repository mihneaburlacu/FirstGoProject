package main

import (
	"fmt"

	"example.com/go-demo-1/filesio"
	"example.com/go-demo-1/parsercsv"
)

func main() {
	fmt.Println("Start project")
	fmt.Println()

	nrOfChunks := 100
	fileName := "input.csv"

	data, err := filesio.ReturnData(fileName)

	if err != nil {
		fmt.Println("Error reading data from file")
		return
	}

	fmt.Println("ID  first_name  last_name  email  gender  ip_address")
	fmt.Println()

	allPersonalDetails := parsercsv.Parser(data, nrOfChunks)
	err = filesio.WriteFiles(allPersonalDetails)

	if err != nil {
		fmt.Println("Error writing data in file")
		return
	}

	fmt.Println()
	fmt.Println("End project")
}
