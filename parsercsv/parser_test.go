package parsercsv

import (
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
