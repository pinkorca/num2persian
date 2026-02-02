package num2persian

import (
	"math"
	"math/big"
	"testing"
)

func TestConvert_BasicNumbers(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "صفر"},
		{1, "یک"},
		{2, "دو"},
		{3, "سه"},
		{4, "چهار"},
		{5, "پنج"},
		{6, "شش"},
		{7, "هفت"},
		{8, "هشت"},
		{9, "نه"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Teens(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{10, "ده"},
		{11, "یازده"},
		{12, "دوازده"},
		{13, "سیزده"},
		{14, "چهارده"},
		{15, "پانزده"},
		{16, "شانزده"},
		{17, "هفده"},
		{18, "هجده"},
		{19, "نوزده"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Tens(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{20, "بیست"},
		{30, "سی"},
		{40, "چهل"},
		{50, "پنجاه"},
		{60, "شصت"},
		{70, "هفتاد"},
		{80, "هشتاد"},
		{90, "نود"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_CompoundTens(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{21, "بیست و یک"},
		{32, "سی و دو"},
		{45, "چهل و پنج"},
		{58, "پنجاه و هشت"},
		{63, "شصت و سه"},
		{77, "هفتاد و هفت"},
		{84, "هشتاد و چهار"},
		{99, "نود و نه"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Hundreds(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{100, "صد"},
		{200, "دویست"},
		{300, "سیصد"},
		{400, "چهارصد"},
		{500, "پانصد"},
		{600, "ششصد"},
		{700, "هفتصد"},
		{800, "هشتصد"},
		{900, "نهصد"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_CompoundHundreds(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{101, "صد و یک"},
		{110, "صد و ده"},
		{111, "صد و یازده"},
		{125, "صد و بیست و پنج"},
		{250, "دویست و پنجاه"},
		{365, "سیصد و شصت و پنج"},
		{499, "چهارصد و نود و نه"},
		{512, "پانصد و دوازده"},
		{789, "هفتصد و هشتاد و نه"},
		{999, "نهصد و نود و نه"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Thousands(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1000, "هزار"},
		{2000, "دو هزار"},
		{5000, "پنج هزار"},
		{10000, "ده هزار"},
		{15000, "پانزده هزار"},
		{100000, "صد هزار"},
		{500000, "پانصد هزار"},
		{999000, "نهصد و نود و نه هزار"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_CompoundThousands(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1001, "هزار و یک"},
		{1010, "هزار و ده"},
		{1100, "هزار و صد"},
		{1234, "هزار و دویست و سی و چهار"},
		{5678, "پنج هزار و ششصد و هفتاد و هشت"},
		{12345, "دوازده هزار و سیصد و چهل و پنج"},
		{99999, "نود و نه هزار و نهصد و نود و نه"},
		{100001, "صد هزار و یک"},
		{123456, "صد و بیست و سه هزار و چهارصد و پنجاه و شش"},
		{999999, "نهصد و نود و نه هزار و نهصد و نود و نه"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Millions(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1000000, "یک میلیون"},
		{2000000, "دو میلیون"},
		{10000000, "ده میلیون"},
		{100000000, "صد میلیون"},
		{999000000, "نهصد و نود و نه میلیون"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_CompoundMillions(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1000001, "یک میلیون و یک"},
		{1001000, "یک میلیون و هزار"},
		{1234567, "یک میلیون و دویست و سی و چهار هزار و پانصد و شصت و هفت"},
		{999999999, "نهصد و نود و نه میلیون و نهصد و نود و نه هزار و نهصد و نود و نه"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Billions(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1000000000, "یک میلیارد"},
		{2000000000, "دو میلیارد"},
		{10000000000, "ده میلیارد"},
		{100000000000, "صد میلیارد"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Trillions(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1000000000000, "یک تریلیون"},
		{5000000000000, "پنج تریلیون"},
		{100000000000000, "صد تریلیون"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Quadrillions(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1000000000000000, "یک کوادریلیون"},
		{5000000000000000, "پنج کوادریلیون"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_Quintillions(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1000000000000000000, "یک کوینتیلیون"},
		{2000000000000000000, "دو کوینتیلیون"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_NegativeNumbers(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{-1, "منفی یک"},
		{-10, "منفی ده"},
		{-100, "منفی صد"},
		{-1000, "منفی هزار"},
		{-1234, "منفی هزار و دویست و سی و چهار"},
		{-999999, "منفی نهصد و نود و نه هزار و نهصد و نود و نه"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{"MaxInt64", math.MaxInt64, "نه کوینتیلیون و دویست و بیست و سه کوادریلیون و سیصد و هفتاد و دو تریلیون و سی و شش میلیارد و هشتصد و پنجاه و چهار میلیون و هفتصد و هفتاد و پنج هزار و هشتصد و هفت"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertBigInt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0", "صفر"},
		{"1", "یک"},
		// Sextillion (10^21)
		{"1000000000000000000000", "یک سکستیلیون"},
		{"-1000000000000000000000", "منفی یک سکستیلیون"},
		// Septillion (10^24)
		{"1000000000000000000000000", "یک سپتیلیون"},
		{"5000000000000000000000000", "پنج سپتیلیون"},
		// Octillion (10^27)
		{"1000000000000000000000000000", "یک اکتیلیون"},
		{"7000000000000000000000000000", "هفت اکتیلیون"},
		// Nonillion (10^30)
		{"1000000000000000000000000000000", "یک نونیلیون"},
		{"3000000000000000000000000000000", "سه نونیلیون"},
		// Decillion (10^33)
		{"1000000000000000000000000000000000", "یک دسیلیون"},
		{"9000000000000000000000000000000000", "نه دسیلیون"},
		// Compound large number
		{"1001000000000000000000000000000000", "یک دسیلیون و یک نونیلیون"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			n := new(big.Int)
			n.SetString(tt.input, 10)
			result := ConvertBigInt(n)
			if result != tt.expected {
				t.Errorf("ConvertBigInt(%s) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertBigInt_Nil(t *testing.T) {
	result := ConvertBigInt(nil)
	if result != "صفر" {
		t.Errorf("ConvertBigInt(nil) = %q, want %q", result, "صفر")
	}
}

func TestConvertFloat(t *testing.T) {
	tests := []struct {
		input     float64
		precision int
		expected  string
	}{
		{0.0, 0, "صفر"},
		{0.0, 1, "صفر ممیز صفر"},
		{1.5, 1, "یک ممیز پنج"},
		{3.14, 2, "سه ممیز چهارده"},
		{123.456, 3, "صد و بیست و سه ممیز چهارصد و پنجاه و شش"},
		{-12.5, 1, "منفی دوازده ممیز پنج"},
		{99.99, 2, "نود و نه ممیز نود و نه"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := ConvertFloat(tt.input, tt.precision)
			if result != tt.expected {
				t.Errorf("ConvertFloat(%f, %d) = %q, want %q", tt.input, tt.precision, result, tt.expected)
			}
		})
	}
}

func TestConvertFloat_SpecialValues(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"NaN", math.NaN(), "نامعین"},
		{"PositiveInfinity", math.Inf(1), "بی‌نهایت"},
		{"NegativeInfinity", math.Inf(-1), "منفی بی‌نهایت"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertFloat(tt.input, 2)
			if result != tt.expected {
				t.Errorf("ConvertFloat(%v, 2) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"0", "صفر", false},
		{"123", "صد و بیست و سه", false},
		{"-456", "منفی چهارصد و پنجاه و شش", false},
		{"  789  ", "هفتصد و هشتاد و نه", false},
		{"12.5", "دوازده ممیز پنج", false},
		{"abc", "", true},
		{"12.34.56", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ConvertString(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ConvertString(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("ConvertString(%q) unexpected error: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("ConvertString(%q) = %q, want %q", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestConvertInt(t *testing.T) {
	result := ConvertInt(42)
	expected := "چهل و دو"
	if result != expected {
		t.Errorf("ConvertInt(42) = %q, want %q", result, expected)
	}
}

func TestParseError(t *testing.T) {
	err := &ParseError{Input: "test"}
	expected := `num2persian: cannot parse "test" as a number`
	if err.Error() != expected {
		t.Errorf("ParseError.Error() = %q, want %q", err.Error(), expected)
	}
}

// Benchmark tests
func BenchmarkConvert_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(42)
	}
}

func BenchmarkConvert_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(1234567)
	}
}

func BenchmarkConvert_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(9223372036854775807)
	}
}

func BenchmarkConvertBigInt(b *testing.B) {
	n := new(big.Int)
	n.SetString("1000000000000000000000000000000000", 10) // decillion
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBigInt(n)
	}
}
