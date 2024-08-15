package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 to I", 1, "I"},
		{"2 to II", 2, "II"},
		{"3 to III", 3, "III"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			assertConversionResult(t, ConvertToRoman(test.Arabic), test.Want)
		})
	}
}

func assertConversionResult(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
