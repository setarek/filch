package filch

func ConvertToPersianDigit(text string) (output string) {

	persianDigitMap := map[string]string{
		"0": "۰",
		"1": "۱",
		"2": "۲",
		"3": "۳",
		"4": "۴",
		"5": "۵",
		"6": "۶",
		"7": "۷",
		"8": "۸",
		"9": "۹",
		"٤": "۴",
		"٥": "۵",
		"٦": "۶",
	}

	for _, char := range text {
		if pch, ok := persianDigitMap[string(char)]; ok {
			output += pch
		} else {
			output += string(char)
		}
	}
	return output
}

func ConvertArabicToPersian(text string) (output string) {

	arabicToPersianMap := map[string]string{
		"أ": "ا",
		"إ": "ا",
		"ك": "ک",
		"ي": "ی",
		"ئ": "ی",
	}

	for _, char := range text {
		if pch, ok := arabicToPersianMap[string(char)]; ok {
			output += pch
		} else {
			output += string(char)
		}
	}
	return output
}