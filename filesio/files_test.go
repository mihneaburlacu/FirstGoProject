package filesio

import (
	"testing"

	"example.com/go-demo-1/parsercsv"
)

func TestReturnData(t *testing.T) {
	var want [][]string
	want = append(want, []string{"id", "first_name", "last_name", "email", "gender", "ip_address"})
	want = append(want, []string{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"})

	got, err := ReturnData("test.csv")

	if err != nil {
		t.Errorf("Error")
	}

	theyAreEquals := true
	for i, line := range want {
		for j, _ := range line {
			if want[i][j] != got[i][j] {
				theyAreEquals = false
			}
		}
	}

	if !theyAreEquals {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestWriteFiles(t *testing.T) {
	var allPersonalDetails [][]parsercsv.PersonalDetailsRecord
	var oneChunkOfPersonalDetails []parsercsv.PersonalDetailsRecord
	var onePerson parsercsv.PersonalDetailsRecord

	parsercsv.Load(&onePerson, []string{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"})
	oneChunkOfPersonalDetails = append(oneChunkOfPersonalDetails, onePerson)
	parsercsv.Load(&onePerson, []string{"2", "Alex", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"})
	oneChunkOfPersonalDetails = append(oneChunkOfPersonalDetails, onePerson)

	allPersonalDetails = append(allPersonalDetails, oneChunkOfPersonalDetails)

	got, err := WriteFiles(allPersonalDetails)
	want := "id,first_name,last_name,email,gender,ip_address\n1 Mavra Malec mmalec0@usa.gov Female 229.215.245.102\n2 Alex Malec mmalec0@usa.gov Female 229.215.245.102"

	if got[0] != want && err != nil {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
