package CalculatorService

type CalculatorService interface {
	Sum(number1 int, number2 int) int
	SumALotOfArgs(numbers ...int) (result int)
}

type calculatorService struct {
}

func NewCalculatorService() (c CalculatorService) {
	return &calculatorService{}
}

func (c *calculatorService) Sum(number1 int, number2 int) int {

	return number1 + number2
}

func (c *calculatorService) SumALotOfArgs(numbers ...int) (result int) {
	result = 0
	for _, number := range numbers {
		result += number
	}

	return
}
