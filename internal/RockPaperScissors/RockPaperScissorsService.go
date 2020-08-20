package RockPaperScissors

type Base interface {
	IsTheWinner(x Base) bool
}

type base struct {
	Win []Base
}

func (b *base) IsTheWinner(x Base) bool {
	for _, n := range b.Win {
		if x == n {
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
			[]Base{NewScissor()},
		},
	}
}

type Paper struct {
	*base
}

func NewPaper() (o Paper) {
	return Paper{
		&base{
			[]Base{NewRock()},
		},
	}
}

type Scissor struct {
	*base
}

func NewScissor() (o Scissor) {
	return Scissor{
		&base{
			[]Base{NewPaper()},
		},
	}
}

type RockPaperScissorsService interface {
	Play(player1 Base, player2 Base) (b Base)
}

type rockPaperScissorsService struct {
}

func NewRockPaperScissorsService() (o RockPaperScissorsService) {
	return &rockPaperScissorsService{}
}

func (o *rockPaperScissorsService) Play(player1 Base, player2 Base) (b Base) {
	if player1 == player2 {
		b = player1
	}

	if player1.IsTheWinner(player2) {
		return player1
	}

	return player2
}
