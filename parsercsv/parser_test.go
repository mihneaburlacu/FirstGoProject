package parsercsv

import (
	"reflect"
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
	type testCases struct {
		name  string
		input [][]string
		want  [][]PersonalDetailsRecord
	}

	for _, scenario := range []testCases{
		{
			name:  "all chunks are ok",
			input: [][]string{{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"}, {"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"}, {"3", "Gerri", "Choffin", "gchoffin2@ning.com", "Male", "9.254.198.50"}, {"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"}},
			want: [][]PersonalDetailsRecord{
				{
					PersonalDetailsRecord{
						ID:        "1",
						FirstName: "Mavra",
						LastName:  "Malec",
						Email:     "mmalec0@usa.gov",
						Gender:    "Female",
						IpAddress: "229.215.245.102",
					},
					PersonalDetailsRecord{
						ID:        "2",
						FirstName: "Fan",
						LastName:  "Gilvear",
						Email:     "fgilvear1@people.com.cn",
						Gender:    "Female",
						IpAddress: "125.219.253.132",
					},
				},
				{
					PersonalDetailsRecord{
						ID:        "3",
						FirstName: "Gerri",
						LastName:  "Choffin",
						Email:     "gchoffin2@ning.com",
						Gender:    "Male",
						IpAddress: "9.254.198.50",
					},
					PersonalDetailsRecord{
						ID:        "4",
						FirstName: "Tremayne",
						LastName:  "Loosemore",
						Email:     "tloosemore3@cnn.com",
						Gender:    "Male",
						IpAddress: "167.249.115.222",
					},
				},
			},
		},
		{
			name:  "one field is missing",
			input: [][]string{{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"}, {"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"}, {"3", "Gerri", "Choffin", "gchoffin2@ning.com", "Male", "9.254.198.50"}, {"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "167.249.115.222"}},
			want: [][]PersonalDetailsRecord{
				{
					PersonalDetailsRecord{
						ID:        "1",
						FirstName: "Mavra",
						LastName:  "Malec",
						Email:     "mmalec0@usa.gov",
						Gender:    "Female",
						IpAddress: "229.215.245.102",
					},
					PersonalDetailsRecord{
						ID:        "2",
						FirstName: "Fan",
						LastName:  "Gilvear",
						Email:     "fgilvear1@people.com.cn",
						Gender:    "Female",
						IpAddress: "125.219.253.132",
					},
				},
				{
					PersonalDetailsRecord{
						ID:        "3",
						FirstName: "Gerri",
						LastName:  "Choffin",
						Email:     "gchoffin2@ning.com",
						Gender:    "Male",
						IpAddress: "9.254.198.50",
					},
				},
			},
		},
		{
			name:  "the input data has a header",
			input: [][]string{{"id", "first_name", "last_name", "email", "gender", "ip_address"}, {"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"}},
			want: [][]PersonalDetailsRecord{
				{
					PersonalDetailsRecord{
						ID:        "1",
						FirstName: "Mavra",
						LastName:  "Malec",
						Email:     "mmalec0@usa.gov",
						Gender:    "Female",
						IpAddress: "229.215.245.102",
					},
				},
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			got := Parser(scenario.input, 2)
			if !reflect.DeepEqual(got, scenario.want) {
				t.Errorf("got %v, wanted %v", got, scenario.want)
			}
		})
	}

}
