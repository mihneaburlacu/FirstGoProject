package parserCSVFile

type PersonalDetailsRecord struct {
	ID        string
	firstName string
	lastName  string
	email     string
	gender    string
	ipAddress string
}

func GetID(record PersonalDetailsRecord) string {
	return record.ID
}

func GetFirstName(record PersonalDetailsRecord) string {
	return record.firstName
}

func GetLastName(record PersonalDetailsRecord) string {
	return record.lastName
}

func GetEmail(record PersonalDetailsRecord) string {
	return record.email
}

func GetGender(record PersonalDetailsRecord) string {
	return record.gender
}

func GetIpAddress(record PersonalDetailsRecord) string {
	return record.ipAddress
}

func GetAllDetails(record PersonalDetailsRecord) string {
	return record.ID + " " + record.firstName + " " + record.lastName + " " + record.email + " " + record.gender + " " + record.ipAddress
}

func Parser(data [][]string, X int) [][]PersonalDetailsRecord {
	var personalDetails [][]PersonalDetailsRecord

	index := 0
	var chunkOfPersonalDetails []PersonalDetailsRecord

	for i, line := range data {
		if i > 0 {
			var onePerson PersonalDetailsRecord
			nrFields := 0

			for j, field := range line {
				if j == 0 {
					onePerson.ID = field
					if field != "" {
						nrFields++
					}
				} else if j == 1 {
					onePerson.firstName = field
					if field != "" {
						nrFields++
					}
				} else if j == 2 {
					onePerson.lastName = field
					if field != "" {
						nrFields++
					}
				} else if j == 3 {
					onePerson.email = field
					if field != "" {
						nrFields++
					}
				} else if j == 4 {
					onePerson.gender = field
					if field != "" {
						nrFields++
					}
				} else {
					onePerson.ipAddress = field
					if field != "" {
						nrFields++
					}
				}
			}

			if nrFields == 6 {
				chunkOfPersonalDetails = append(chunkOfPersonalDetails, onePerson)
				index = index + 1
			}

			if index == X || i == len(data)-1 {
				index = 0
				personalDetails = append(personalDetails, chunkOfPersonalDetails)
				var emptyStruct []PersonalDetailsRecord
				chunkOfPersonalDetails = emptyStruct
			}

		}
	}

	return personalDetails
}
