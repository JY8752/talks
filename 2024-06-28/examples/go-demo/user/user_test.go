package user_test

import "testing"

func TestExample(t *testing.T) {
	tests := map[string]struct {
		x        int
		y        int
		expected int
	}{
		"1 + 1 = 2": {
			x:        1,
			y:        1,
			expected: 2,
		},
		"2 + 3 = 5": {
			x:        2,
			y:        3,
			expected: 5,
		},
		"0 + (-1) = -1": {
			x:        0,
			y:        -1,
			expected: -1,
		},
	}

	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			if tt.x+tt.y != tt.expected {
				t.Errorf("expected %d but %d", tt.expected, tt.x+tt.y)
			}
		})
	}
}
