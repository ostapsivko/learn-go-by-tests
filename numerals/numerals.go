package numerals

import "strings"

func ConvertToRoman(value int) string {
	var result strings.Builder

	for value > 0 {
		switch {
		case value > 4:
			result.WriteString("V")
			value -= 5
		case value > 3:
			result.WriteString("IV")
			value -= 4
		default:
			result.WriteString("I")
			value--
		}
	}

	return result.String()
}
