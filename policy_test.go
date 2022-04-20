package mdschema

import (
	"fmt"
	"testing"
)

func Test_IntPolicyValidation(t *testing.T) {
	p := IntPolicy{
		Required: true,
		Start:    0,
		Stop:     10,
		Step:     1,
	}

	tests := []struct {
		input    int
		expected bool
	}{
		{0, true},
		{2, true},
		{6, true},
		{9, true},
		{10, true},
		{11, false},
		{-1, false},
		{-100, false},
		{-512, false},
	}

	t.Run("Range(0,10,1)", func(t *testing.T) {
		for _, tc := range tests {
			t.Run(fmt.Sprintf("%d", tc.input), func(t *testing.T) {
				if p.Validate(tc.input) != tc.expected {
					t.Logf("Validating %d, expected %t\n", tc.input, p.Validate(tc.input))
					t.Fail()
				}
			})
		}
	})
}
