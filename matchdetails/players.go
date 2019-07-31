package matchdetails

import (
	"fmt"
	"strconv"

	"github.com/jmcvetta/randutil"
)

// Player struct stores the details of a player
type Player struct {
	Name           string
	CurrentScore   int
	IsPlaying      bool
	OnStrike       bool
	BallsPlayed    int
	ShotProbablity []randutil.Choice
}

// UpdateDetails is used to mutate values of a player
func (player *Player) UpdateDetails(incrementscoreby int, isOut bool) {
	if isOut {
		player.IsPlaying = false
	}
	player.CurrentScore += incrementscoreby
	player.BallsPlayed++
}

//ChangeStrike would be called after a player hits 1, 3, 5 or if over changes
func (player *Player) ChangeStrike() {
	player.OnStrike = !player.OnStrike
}

//ScoreCard is used at the end of match to display player runs and balls played
func ScoreCard(players []Player) {
	for _, player := range players {
		status := "*"
		if player.IsPlaying {
			status = "*"
		} else {
			status = " "
		}
		resstr := player.Name + status + "  " + strconv.Itoa(player.CurrentScore) + "(" + strconv.Itoa(player.BallsPlayed) + ")"
		fmt.Println(resstr)
	}
}
