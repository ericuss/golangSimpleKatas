package Calculator

type Calculator struct {
	Name string
}

func NewCalculator(name string) (c Calculator) {
	c = Calculator{
		Name: name,
	}
	return
}
