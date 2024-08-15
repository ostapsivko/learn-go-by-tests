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
		{"40 to XL", 40, "XL"},
		{"48 to XLVIII", 48, "XLVIII"},
		{"49 to XLIX", 49, "XLIX"},
		{"50 to L", 50, "L"},
		{"54 to LIV", 54, "LIV"},
		{"55 to LV", 55, "LV"},
		{"59 to LIX", 59, "LIX"},
		{"60 to LX", 60, "LX"},
		{"64 to LXIV", 64, "LXIV"},
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
