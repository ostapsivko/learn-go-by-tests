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
		{"4 to IV", 4, "IV"},
		{"5 to V", 5, "V"},
		{"6 to VI", 6, "VI"},
		{"7 to VII", 7, "VII"},
		{"8 to VIII", 8, "VIII"},
		{"9 to IX", 9, "IX"},
		{"10 to X", 10, "X"},
		{"11 to XI", 11, "XI"},
		{"12 to XII", 12, "XII"},
		{"14 to XIV", 14, "XIV"},
		{"15 to XV", 15, "XV"},
		{"18 to XVIII", 18, "XVIII"},
		{"20 to XX", 20, "XX"},
		{"25 to XXV", 25, "XXV"},
		{"39 to XXXIX", 39, "XXXIX"},
		{"44 to XXXXIV", 44, "XXXXIV"},
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
