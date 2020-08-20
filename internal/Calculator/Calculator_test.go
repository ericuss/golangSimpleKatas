package Calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorIsConstructed(t *testing.T) {
	name := "Bucefalo"
	c := NewCalculator(name)

	if c.Name == name {
		assert.Error(t, errors.New("error"))
	}

}
