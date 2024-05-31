package utilities

func FormatIntRoman(num int) string {
	// Format a number as lowercase Roman numerals.
	roman := map[int]string{
		1000: "m",
		900:  "cm",
		500:  "d",
		400:  "cd",
		100:  "c",
		90:   "xc",
		50:   "l",
		40:   "xl",
		10:   "x",
		9:    "ix",
		5:    "v",
		4:    "iv",
		1:    "i",
	}
	result := ""
	for value, numeral := range roman {
		for num >= value {
			result += numeral
			num -= value
		}
	}
	return result
}
