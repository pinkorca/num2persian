package num2persian

import "testing"

func TestConvertOrdinal(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1, "اول"},
		{2, "دوم"},
		{3, "سوم"},
		{4, "چهارم"},
		{5, "پنجم"},
		{10, "دهم"},
		{11, "یازدهم"},
		{21, "بیست و یکم"},
		{23, "بیست و سوم"},
		{100, "صدم"},
		{103, "صد و سوم"},
		{1000, "هزارم"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := ConvertOrdinal(tt.input)
			if result != tt.expected {
				t.Errorf("ConvertOrdinal(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertOrdinal_Zero(t *testing.T) {
	result := ConvertOrdinal(0)
	if result != "" {
		t.Errorf("ConvertOrdinal(0) = %q, want empty string", result)
	}
}

func TestConvertOrdinal_Negative(t *testing.T) {
	result := ConvertOrdinal(-1)
	if result != "" {
		t.Errorf("ConvertOrdinal(-1) = %q, want empty string", result)
	}
}

func TestConvertOrdinalInt(t *testing.T) {
	result := ConvertOrdinalInt(5)
	expected := "پنجم"
	if result != expected {
		t.Errorf("ConvertOrdinalInt(5) = %q, want %q", result, expected)
	}
}
