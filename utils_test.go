package envenc

import (
	"testing"
)

func TestStrRandom(t *testing.T) {
	tests := []struct {
		name           string
		length         int
		expectedLength int
	}{
		{"Test length 10", 10, 10},
		{"Test length 20", 20, 20},
		{"Test length 0", 0, 0},
		{"Test length -1", -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := strRandom(tt.length)

			if tt.length == 0 || tt.expectedLength == 0 {
				if str != "" {
					t.Errorf("expected empty string, got %q", str)
				}
				return
			}

			if len(str) != tt.expectedLength {
				t.Errorf("expected string of length %d, got %q", tt.expectedLength, str)
			}

			// Test that the function returns a different string each time it is called
			str2 := strRandom(tt.length)
			if str == str2 {
				t.Errorf("expected different strings, got %q and %q", str, str2)
			}
		})
	}
}
