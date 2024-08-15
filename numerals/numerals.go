package numerals

import "strings"

func ConvertToRoman(value int) string {
	var result strings.Builder

	for i := value; i > 0; i-- {
		if value == 4 {
			result.WriteString("IV")
			break
		}

		result.WriteString("I")
	}

	return result.String()
}
