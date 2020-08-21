package RockPaperScissors

import "reflect"

type Base interface {
	IsTheWinner(x Base) bool
}

type base struct {
	Win []Base
}

func (b *base) IsTheWinner(x Base) bool {
	for _, n := range b.Win {
		if reflect.TypeOf(x) == reflect.TypeOf(n) {
			return true
		}
	}
	return false
}

type Rock struct {
	*base
}

func NewRock() (o Rock) {
	return Rock{
		&base{
			[]Base{Scissor{}},
		},
	}
}

type Paper struct {
	*base
}

func NewPaper() (o Paper) {
	return Paper{
		&base{
			[]Base{Rock{}},
		},
	}
}

type Scissor struct {
	*base
}

func NewScissor() (o Scissor) {
	return Scissor{
		&base{
			[]Base{Paper{}},
		},
	}
}
