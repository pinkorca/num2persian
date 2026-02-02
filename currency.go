package num2persian

const (
	tomanUnit = "تومان"
	rialUnit  = "ریال"
)

// ToToman converts a number to Persian text with "تومان" suffix.
func ToToman(n int64) string {
	return Convert(n) + " " + tomanUnit
}

// ToTomanInt converts an int to Persian text with "تومان" suffix.
func ToTomanInt(n int) string {
	return ToToman(int64(n))
}

// ToRial converts a number to Persian text with "ریال" suffix.
func ToRial(n int64) string {
	return Convert(n) + " " + rialUnit
}

// ToRialInt converts an int to Persian text with "ریال" suffix.
func ToRialInt(n int) string {
	return ToRial(int64(n))
}

// TomanToRial converts Toman to Rial and returns Persian text.
func TomanToRial(n int64) string {
	return ToRial(n * 10)
}

// RialToToman converts Rial to Toman and returns Persian text.
func RialToToman(n int64) string {
	return ToToman(n / 10)
}
