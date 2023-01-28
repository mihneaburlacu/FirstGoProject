package parsercsv

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	gotTrue := Validate(PersonalDetailsRecord{"10", "Mihnea", "Burlacu", "mihnea@gmail.com", "Male", "115.200.300"})
	wantTrue := true

	if gotTrue != wantTrue {
		t.Errorf("got %t, wanted %t", gotTrue, wantTrue)
	}

	gotFalse := Validate(PersonalDetailsRecord{"15", "Calin", "Burlacu", "calin@gmail.com", "", "115.200.300"})
	wantFalse := false

	if gotFalse != wantFalse {
		t.Errorf("got %t, wanted %t", gotFalse, wantFalse)
	}
}

func TestLoad(t *testing.T) {
	var onePerson PersonalDetailsRecord
	data := [6]string{"10", "Mihnea", "Burlacu", "mihnea@gmail.com", "Male", "115.200.300"}
	Load(&onePerson, data[:])

	secondPerson := PersonalDetailsRecord{"10", "Mihnea", "Burlacu", "mihnea@gmail.com", "Male", "115.200.300"}

	if onePerson != secondPerson {
		t.Errorf("wanted the first and second person to have the same fields, but they did not")
	}

	thirdPerson := PersonalDetailsRecord{"15", "Calin", "Burlacu", "calin@gmail.com", "", "115.200.300"}

	if onePerson == thirdPerson {
		t.Errorf("wanted the first and the third person to have not the same fields, but they had")
	}
}

func TestShowAllDetails(t *testing.T) {
	got := ShowAllDetails(PersonalDetailsRecord{"10", "Mihnea", "Burlacu", "mihnea@gmail.com", "Male", "115.200.300"})
	want := "10 Mihnea Burlacu mihnea@gmail.com Male 115.200.300"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestParser(t *testing.T) {
	var data [][]string
	data = append(data, []string{"id", "first_name", "last_name", "email", "gender", "ip_address"})
	data = append(data, []string{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"})
	data = append(data, []string{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"})
	data = append(data, []string{"3", "Gerri", "Choffin", "gchoffin2@ning.com", "", "9.254.198.50"})
	data = append(data, []string{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"})

	var want [][]PersonalDetailsRecord
	var oneChunkOfPersonalDetails []PersonalDetailsRecord
	var onePerson PersonalDetailsRecord
	Load(&onePerson, []string{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"})
	oneChunkOfPersonalDetails = append(oneChunkOfPersonalDetails, onePerson)
	Load(&onePerson, []string{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"})
	oneChunkOfPersonalDetails = append(oneChunkOfPersonalDetails, onePerson)
	want = append(want, oneChunkOfPersonalDetails)
	oneChunkOfPersonalDetails = nil
	Load(&onePerson, []string{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"})
	oneChunkOfPersonalDetails = append(oneChunkOfPersonalDetails, onePerson)
	want = append(want, oneChunkOfPersonalDetails)

	got := Parser(data, 2)

	if len(got) != len(want) {
		t.Errorf("The length does not match")
	}

	theyAreEquals := true
	for i, oneChunk := range want {
		for j, _ := range oneChunk {
			if got[i][j] != want[i][j] {
				fmt.Println(got[i][j])
				fmt.Println(want[i][j])
				theyAreEquals = false
			}
		}
	}

	if theyAreEquals == false {
		t.Errorf("they are not equal")
	}
}
