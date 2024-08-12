package walker

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Azdab", 27},
			[]string{"Azdab"},
		},
		{
			"struct with nested fields",
			Person{
				"Azdab", Profile{27, "Mykolaiv"},
			},
			[]string{"Azdab", "Mykolaiv"},
		},
		{
			"pointers to things",
			&Person{
				"Azdab",
				Profile{27, "Mykolaiv"},
			},
			[]string{"Azdab", "Mykolaiv"},
		},
		{
			"slices",
			[]Person{
				{"Azdab", Profile{27, "Mykolaiv"}},
				{"Not Azdab", Profile{27, "Odesa"}},
			},
			[]string{"Azdab", "Mykolaiv", "Not Azdab", "Odesa"},
		},
		{
			"arrays",
			[1]Person{
				{"Azdab", Profile{27, "Mykolaiv"}},
			},
			[]string{"Azdab", "Mykolaiv"},
		},
		{
			"maps",
			map[string]Person{
				"Azdab": {"Azdab", Profile{27, "Mykolaiv"}},
			},
			[]string{"Azdab", "Mykolaiv"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, test.ExpectedCalls)
			}

		})
	}
}
