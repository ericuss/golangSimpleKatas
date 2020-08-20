package RockPaperScissors

import (
	"fmt"
	"reflect"
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
		"Sum Rock + Scissor":    {player1: NewRock(), player2: NewScissor(), want: Rock{}},
		"Sum Scissor + Scissor": {player1: NewScissor(), player2: NewScissor(), want: NewScissor()},
		"Sum Scissor + Rock":    {player1: NewScissor(), player2: NewRock(), want: NewRock()},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := NewRockPaperScissorsService()
			playerResult := service.Play(tc.player1, tc.player2)
			fmt.Print(reflect.TypeOf(tc.want))
			fmt.Print(reflect.TypeOf(playerResult))
			assert.Equal(t, reflect.TypeOf(tc.want), reflect.TypeOf(playerResult))
		})
	}
}
