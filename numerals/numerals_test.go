package numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Description string
	Arabic      uint16
	Roman       string
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
	{"80 to LXXX", 80, "LXXX"},
	{"89 to LXXXIX", 89, "LXXXIX"},
	{"90 to XC", 90, "XC"},
	{"95 to XCV", 95, "XCV"},
	{"99 to XCIX", 99, "XCIX"},
	{"100 to C", 100, "C"},
	{"188 to CLXXXVIII", 188, "CLXXXVIII"},
	{"154 to CLIV", 154, "CLIV"},
	{"199 to CXCIX", 199, "CXCIX"},
	{"200 to CC", 200, "CC"},
	{"400 to CD", 400, "CD"},
	{"500 to D", 500, "D"},
	{"900 to CM", 900, "CM"},
	{"1000 to M", 1000, "M"},
	{"1500 to MD", 1500, "MD"},
	{"1984 to MCMLXXXIV", 1984, "MCMLXXXIV"},
	{"1997 to MCMXCVII", 1997, "MCMXCVII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			assertConversionResult(t, ConvertToRoman(test.Arabic), test.Roman)
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %q", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestConvertingToArabicRec(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %q", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabicRec(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func BenchmarkConvertingToArabic(b *testing.B) {
	testCase := cases[len(cases)-1]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConvertToArabic(testCase.Roman)
	}
}

func BenchmarkConvertingToArabicRec(b *testing.B) {
	testCase := cases[len(cases)-1]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertToArabicRec(testCase.Roman)
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

func assertConversionResult(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
