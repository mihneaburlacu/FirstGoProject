package parsercsv

type PersonalDetailsRecord struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Gender    string
	IpAddress string
}

func Validate(record PersonalDetailsRecord) bool {
	return !(record.ID == "" || record.FirstName == "" || record.LastName == "" || record.Email == "" || record.Gender == "" || record.IpAddress == "")
}

func Load(record *PersonalDetailsRecord, line []string) {
	if len(line) == 6 {
		*&record.ID = line[0]
		*&record.FirstName = line[1]
		*&record.LastName = line[2]
		*&record.Email = line[3]
		*&record.Gender = line[4]
		*&record.IpAddress = line[5]
	}
}

func ShowAllDetails(record PersonalDetailsRecord) string {
	return record.ID + " " + record.FirstName + " " + record.LastName + " " + record.Email + " " + record.Gender + " " + record.IpAddress
}

func Parser(data [][]string, nrOfChunks int) [][]PersonalDetailsRecord {
	var personalDetails [][]PersonalDetailsRecord

	index := 0
	var chunkOfPersonalDetails []PersonalDetailsRecord

	for i, line := range data {
		if i > 0 {
			var onePerson PersonalDetailsRecord

			Load(&onePerson, line)

			if Validate(onePerson) {
				chunkOfPersonalDetails = append(chunkOfPersonalDetails, onePerson)
				index = index + 1
			}

			if index == nrOfChunks || i == len(data)-1 {
				index = 0
				personalDetails = append(personalDetails, chunkOfPersonalDetails)
				var emptyStruct []PersonalDetailsRecord
				chunkOfPersonalDetails = emptyStruct
			}

		}
	}

	return personalDetails
}
