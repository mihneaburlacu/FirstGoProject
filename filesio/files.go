package filesio

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"example.com/go-demo-1/parsercsv"
)

func ReturnData(fileName string) ([][]string, error) {
	var data [][]string

	fileCSV, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the csv file")
		return data, err
	}

	csvReader := csv.NewReader(fileCSV)
	data, err = csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading the csv file")
		return data, err
	}

	err = fileCSV.Close()
	if err != nil {
		fmt.Println("Error closing the csv file")
		return data, err
	}

	return data, err
}

func WriteFiles(allPersonalDetails [][]parsercsv.PersonalDetailsRecord) ([]string, error) {
	var errToReturn error
	var allOutputText []string

	for i, chunk := range allPersonalDetails {
		fileOutputName := "output" + strconv.Itoa(i) + ".txt"
		fileOut, err := os.Create(fileOutputName)

		if err != nil {
			errToReturn = err
			fmt.Println("Error creating the output file " + strconv.Itoa(i))
			return allOutputText, err
		}

		_, err = fileOut.WriteString("id,first_name,last_name,email,gender,ip_address\n")

		if err != nil {
			errToReturn = err
			fmt.Println("Error writing in file " + strconv.Itoa(i))
			return allOutputText, err
		}

		var text string

		for _, onePerson := range chunk {
			text = text + parsercsv.ShowAllDetails(onePerson) + "\n"
		}

		text = text + "\n"
		fmt.Println(text)
		fmt.Println()

		_, err = fileOut.WriteString(text)
		allOutputText = append(allOutputText, "id,first_name,last_name,email,gender,ip_address\n"+text)

		if err != nil {
			errToReturn = err
			fmt.Println("Error writing in file " + strconv.Itoa(i))
			return allOutputText, err
		}

		err = fileOut.Close()
		if err != nil {
			errToReturn = err
			fmt.Println("Error closing the output file " + strconv.Itoa(i))
			return allOutputText, err
		}
	}

	return allOutputText, errToReturn
}
