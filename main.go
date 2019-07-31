package main

import (
	"fmt"

	"github.com/ShramikPG/CricketSimulationinGo/matchdetails"
	"github.com/ShramikPG/CricketSimulationinGo/simulator"
	"github.com/jmcvetta/randutil"
	"github.com/logrusorgru/aurora"
)

const (
	totalDel      = 24
	scoreToBeat   = 40
	startingScore = 160
)

func main() {
	kiratBoli := matchdetails.Player{
		Name:           "Kirat Boli",
		CurrentScore:   0,
		IsPlaying:      true,
		OnStrike:       true,
		BallsPlayed:    0,
		ShotProbablity: []randutil.Choice{{5, 0}, {30, 1}, {25, 2}, {10, 3}, {15, 4}, {1, 5}, {9, 6}, {5, "Out"}},
	}
	nsNodhi := matchdetails.Player{
		Name:           "NS Nodhi",
		CurrentScore:   0,
		IsPlaying:      true,
		OnStrike:       false,
		BallsPlayed:    0,
		ShotProbablity: []randutil.Choice{{10, 0}, {40, 1}, {20, 2}, {5, 3}, {10, 4}, {1, 5}, {4, 6}, {10, "Out"}},
	}
	rRumrah := matchdetails.Player{
		Name:           "R Rumrah",
		CurrentScore:   0,
		IsPlaying:      false,
		OnStrike:       false,
		BallsPlayed:    0,
		ShotProbablity: []randutil.Choice{{20, 0}, {30, 1}, {15, 2}, {5, 3}, {5, 4}, {1, 5}, {4, 6}, {20, "Out"}}, // make out a constant
	}
	shashiHenra := matchdetails.Player{
		Name:           "Shashi Nehra",
		CurrentScore:   0,
		IsPlaying:      false,
		OnStrike:       false,
		BallsPlayed:    0,
		ShotProbablity: []randutil.Choice{{30, 0}, {25, 1}, {5, 2}, {0, 3}, {5, 4}, {1, 5}, {4, 6}, {30, "Out"}},
	}
	matchScore := matchdetails.MatchScore{
		WicketsLeft:    3,
		RunsLeft:       scoreToBeat,
		DeliveriesLeft: 24,
	}
	currentPlayers := [2]*matchdetails.Player{&kiratBoli, &nsNodhi}
	playerOnStrike := &matchdetails.Player{}
	playerNotOnStrike := &matchdetails.Player{}
	checkRes := ""

	fmt.Printf("\n LengaBuru is at %v runs, they need 40 more to win with only 3 wickets in hand and 4 overs left \n", startingScore)
	for i := 1; i <= totalDel; i++ {
		checkRes = matchScore.MatchVerdict()
		if checkRes != matchdetails.NotOver {
			fmt.Println(aurora.Bold(checkRes))
			matchdetails.ScoreCard([]matchdetails.Player{kiratBoli, nsNodhi, rRumrah, shashiHenra})
			break
		}
		if currentPlayers[0].OnStrike == true {
			playerOnStrike = currentPlayers[0]
			playerNotOnStrike = currentPlayers[1]
		} else {
			playerOnStrike = currentPlayers[1]
			playerNotOnStrike = currentPlayers[0]
		}
		event, err := randutil.WeightedChoice(playerOnStrike.ShotProbablity)
		if err != nil {
			panic(err)
		}
		simulator.UpdateAfterDelivery(event, playerOnStrike, playerNotOnStrike, &matchScore)
		if event.Item == "Out" {
			simulator.ChangeWickets(currentPlayers, &rRumrah, &shashiHenra, &matchScore)
		}
		if i%6 == 0 {
			playerOnStrike.ChangeStrike()
			playerNotOnStrike.ChangeStrike()
			fmt.Printf("\nThats for this over. The Score is %v-%v(%v). LengaBuru still need %v Runs\n", aurora.Bold(startingScore-matchScore.RunsLeft), aurora.Bold(10-matchScore.WicketsLeft), aurora.Bold(20-matchScore.DeliveriesLeft/6), aurora.Bold(matchScore.RunsLeft))
		}
	}
}
