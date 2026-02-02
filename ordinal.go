package num2persian

import "strings"

var specialOrdinals = map[int64]string{
	1: "اول",
	3: "سوم",
}

// ConvertOrdinal converts an integer to Persian ordinal text.
func ConvertOrdinal(n int64) string {
	if n <= 0 {
		return ""
	}

	if special, ok := specialOrdinals[n]; ok {
		return special
	}

	cardinal := Convert(n)
	if trimmed := strings.TrimSuffix(cardinal, "سه"); trimmed != cardinal {
		return trimmed + "سوم"
	}
	return cardinal + "م"
}

// ConvertOrdinalInt converts an int to Persian ordinal text.
func ConvertOrdinalInt(n int) string {
	return ConvertOrdinal(int64(n))
}
