package RockPaperScissors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlay(t *testing.T) {

	tests := map[string]struct {
		player1 Base
		player2 Base
		want    Base
		// err   error
	}{
		"Sum Rock + Rock":       {player1: NewRock(), player2: NewRock(), want: NewRock()},
		"Sum Rock + Scissor":    {player1: NewRock(), player2: NewScissor(), want: NewRock()},
		"Sum Scissor + Scissor": {player1: NewScissor(), player2: NewScissor(), want: NewScissor()},
		"Sum Scissor + Rock":    {player1: NewScissor(), player2: NewRock(), want: NewRock()},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := NewRockPaperScissorsService()
			playerResult := service.Play(tc.player1, tc.player2)

			assert.Equal(t, tc.want, playerResult)
		})
	}
}
