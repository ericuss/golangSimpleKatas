package RockPaperScissors

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
		return player1
	}

	if player1.IsTheWinner(player2) {
		return player1
	}

	return player2
}
