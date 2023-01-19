package filesInputOutput

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"example.com/go-demo-1/parserCSVFile"
)

func ReturnData(fileName string) [][]string {
	fileCSV, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the csv file")
		log.Fatal(err)
	}

	defer fileCSV.Close()

	csvReader := csv.NewReader(fileCSV)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading the csv file")
		log.Fatal(err)
	}
	//fmt.Println(data)

	return data
}

func WriteFiles(allPersonalDetails [][]parserCSVFile.PersonalDetailsRecord) {
	for i, chunk := range allPersonalDetails {
		fileOutputName := "output" + strconv.Itoa(i) + ".txt"
		fileOut, err2 := os.Create(fileOutputName)

		if err2 != nil {
			log.Fatal(err2)
		}

		defer fileOut.Close()

		_, err3 := fileOut.WriteString("id,first_name,last_name,email,gender,ip_address\n")

		if err3 != nil {
			log.Fatal(err3)
		}

		text := ""

		for _, onePerson := range chunk {
			text = text + parserCSVFile.GetAllDetails(onePerson) + "\n"
		}

		text = text + "\n"
		fmt.Println(text)
		fmt.Println()

		_, err4 := fileOut.WriteString(text)

		if err4 != nil {
			log.Fatal(err4)
		}
	}
}
