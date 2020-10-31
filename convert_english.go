package filch

func ConvertToEnglishDigit(text string) (output string) {
	englishDigitMap := map[string]string{
		"۰": "0",
		"۱": "1",
		"۲": "2",
		"۳": "3",
		"۴": "4",
		"۵": "5",
		"۶": "6",
		"۷": "7",
		"۸": "8",
		"۹": "9",
		"٤": "4",
		"٥": "5",
		"٦": "6",
	}
	for _, char := range text {
		if pchar, ok := englishDigitMap[string(char)]; ok {
			output += pchar
		} else {
			output += string(char)
		}
	}
	return output
}
