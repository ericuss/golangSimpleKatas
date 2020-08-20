package CalculatorService

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {

	tests := map[string]struct {
		input1 int
		input2 int
		want   int
		// err   error
	}{
		"Sum 1 + 1":   {input1: 1, input2: 1, want: 2},
		"Sum 50 + 50": {input1: 50, input2: 50, want: 100},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := NewCalculatorService()
			sumResult := service.Sum(tc.input1, tc.input2)

			assert.Equal(t, tc.want, sumResult)
		})
	}
}

func TestSumALotOfArgs(t *testing.T) {

	tests := map[string]struct {
		inputs []int
		want   int
		// err   error
	}{
		"Sum 1":   {inputs: []int{1}, want: 1},
		"Sum 1,1": {inputs: []int{1, 1}, want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := NewCalculatorService()
			sumResult := service.SumALotOfArgs(tc.inputs[:]...)

			assert.Equal(t, tc.want, sumResult)
		})
	}
}
