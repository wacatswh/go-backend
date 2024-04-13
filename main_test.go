package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := map[string]struct {
		a      int
		b      int
		result int
	}{
		"case 1": {
			a:      1,
			b:      1,
			result: 2,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := Sum(test.a, test.b)
			assert.Equal(t, test.result, actual, "should be 2")
		})
	}
}
