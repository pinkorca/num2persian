// Package num2persian converts numbers to their Persian (Farsi) text representation.
package num2persian

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

var (
	zero      = "صفر"
	negative  = "منفی"
	separator = " و "

	ones = []string{"", "یک", "دو", "سه", "چهار", "پنج", "شش", "هفت", "هشت", "نه"}

	teens = []string{"ده", "یازده", "دوازده", "سیزده", "چهارده", "پانزده",
		"شانزده", "هفده", "هجده", "نوزده"}

	tens = []string{"", "", "بیست", "سی", "چهل", "پنجاه", "شصت", "هفتاد", "هشتاد", "نود"}

	hundreds = []string{"", "صد", "دویست", "سیصد", "چهارصد", "پانصد",
		"ششصد", "هفتصد", "هشتصد", "نهصد"}

	scales = []string{
		"",
		"هزار",
		"میلیون",
		"میلیارد",
		"تریلیون",
		"کوادریلیون",
		"کوینتیلیون",
		"سکستیلیون",
		"سپتیلیون",
		"اکتیلیون",
		"نونیلیون",
		"دسیلیون",
	}
)

// Convert converts an integer to Persian text.
func Convert(n int64) string {
	if n == 0 {
		return zero
	}
	if n < 0 {
		if n == math.MinInt64 {
			return negative + " " + ConvertBigInt(big.NewInt(n).Abs(big.NewInt(n)))
		}
		return negative + " " + Convert(-n)
	}
	return convertPositive(uint64(n))
}

// ConvertInt converts an int to Persian text.
func ConvertInt(n int) string {
	return Convert(int64(n))
}

// ConvertBigInt converts a big.Int to Persian text.
func ConvertBigInt(n *big.Int) string {
	if n == nil || n.Sign() == 0 {
		return zero
	}
	if n.Sign() < 0 {
		return negative + " " + ConvertBigInt(new(big.Int).Abs(n))
	}
	return convertBigIntPositive(n)
}

// ConvertFloat converts a float64 to Persian text with specified decimal precision.
func ConvertFloat(n float64, precision int) string {
	if precision < 0 {
		precision = 0
	}
	if math.IsNaN(n) {
		return "نامعین"
	}
	if math.IsInf(n, 1) {
		return "بی‌نهایت"
	}
	if math.IsInf(n, -1) {
		return negative + " بی‌نهایت"
	}

	isNegative := n < 0
	if isNegative {
		n = -n
	}

	intPart := int64(n)
	multiplier := math.Pow(10, float64(precision))
	decimalPart := int64(math.Round((n - float64(intPart)) * multiplier))

	if decimalPart >= int64(multiplier) {
		intPart++
		decimalPart = 0
	}

	var result strings.Builder
	if isNegative {
		result.WriteString(negative + " ")
	}
	result.WriteString(Convert(intPart))

	if precision > 0 {
		result.WriteString(" ممیز ")
		if decimalPart == 0 {
			result.WriteString(zero)
		} else {
			result.WriteString(Convert(decimalPart))
		}
	}
	return result.String()
}

// ConvertString parses a string and converts it to Persian text.
func ConvertString(s string) (string, error) {
	s = strings.TrimSpace(s)

	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return Convert(n), nil
	}
	if n, err := strconv.ParseFloat(s, 64); err == nil {
		precision := 0
		if idx := strings.Index(s, "."); idx != -1 {
			precision = len(s) - idx - 1
		}
		return ConvertFloat(n, precision), nil
	}

	n := new(big.Int)
	if _, ok := n.SetString(s, 10); ok {
		return ConvertBigInt(n), nil
	}
	return "", &ParseError{Input: s}
}

// ParseError is returned when ConvertString fails to parse the input.
type ParseError struct {
	Input string
}

func (e *ParseError) Error() string {
	return "num2persian: cannot parse \"" + e.Input + "\" as a number"
}

func formatGroupWithScale(group int, scaleIndex int) string {
	text := convertGroup(group)
	if scaleIndex > 0 && text != "" {
		if scaleIndex == 1 && group == 1 {
			return scales[scaleIndex]
		}
		text += " " + scales[scaleIndex]
	}
	return text
}

func convertPositive(n uint64) string {
	if n == 0 {
		return ""
	}

	var parts []string
	scaleIndex := 0

	for n > 0 && scaleIndex < len(scales) {
		group := int(n % 1000)
		n /= 1000

		if group > 0 {
			parts = append([]string{formatGroupWithScale(group, scaleIndex)}, parts...)
		}
		scaleIndex++
	}
	return strings.Join(parts, separator)
}

func convertBigIntPositive(n *big.Int) string {
	if n.Sign() == 0 {
		return ""
	}

	var parts []string
	scaleIndex := 0
	thousand := big.NewInt(1000)
	remaining := new(big.Int).Set(n)
	group := new(big.Int)

	for remaining.Sign() > 0 && scaleIndex < len(scales) {
		remaining.DivMod(remaining, thousand, group)
		groupVal := int(group.Int64())

		if groupVal > 0 {
			parts = append([]string{formatGroupWithScale(groupVal, scaleIndex)}, parts...)
		}
		scaleIndex++
	}

	if remaining.Sign() > 0 {
		remainingText := convertBigIntPositive(remaining)
		parts = append([]string{remainingText + " " + scales[len(scales)-1]}, parts...)
	}
	return strings.Join(parts, separator)
}

func convertGroup(n int) string {
	if n <= 0 || n > 999 {
		return ""
	}

	var parts []string

	if h := n / 100; h > 0 {
		parts = append(parts, hundreds[h])
	}

	remainder := n % 100
	if remainder > 0 {
		if remainder < 10 {
			parts = append(parts, ones[remainder])
		} else if remainder < 20 {
			parts = append(parts, teens[remainder-10])
		} else {
			t := remainder / 10
			o := remainder % 10
			if o > 0 {
				parts = append(parts, tens[t]+separator+ones[o])
			} else {
				parts = append(parts, tens[t])
			}
		}
	}
	return strings.Join(parts, separator)
}
