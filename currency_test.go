package num2persian

import "testing"

func TestToToman(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "صفر تومان"},
		{1, "یک تومان"},
		{1000, "هزار تومان"},
		{1500000, "یک میلیون و پانصد هزار تومان"},
		{-500, "منفی پانصد تومان"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := ToToman(tt.input)
			if result != tt.expected {
				t.Errorf("ToToman(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToTomanInt(t *testing.T) {
	result := ToTomanInt(1000)
	expected := "هزار تومان"
	if result != expected {
		t.Errorf("ToTomanInt(1000) = %q, want %q", result, expected)
	}
}

func TestToRial(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "صفر ریال"},
		{1, "یک ریال"},
		{10000, "ده هزار ریال"},
		{15000000, "پانزده میلیون ریال"},
		{-5000, "منفی پنج هزار ریال"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := ToRial(tt.input)
			if result != tt.expected {
				t.Errorf("ToRial(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToRialInt(t *testing.T) {
	result := ToRialInt(10000)
	expected := "ده هزار ریال"
	if result != expected {
		t.Errorf("ToRialInt(10000) = %q, want %q", result, expected)
	}
}

func TestTomanToRial(t *testing.T) {
	result := TomanToRial(1000)
	expected := "ده هزار ریال"
	if result != expected {
		t.Errorf("TomanToRial(1000) = %q, want %q", result, expected)
	}
}

func TestRialToToman(t *testing.T) {
	result := RialToToman(10000)
	expected := "هزار تومان"
	if result != expected {
		t.Errorf("RialToToman(10000) = %q, want %q", result, expected)
	}
}
