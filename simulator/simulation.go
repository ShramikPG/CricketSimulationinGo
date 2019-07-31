package simulator

import (
	"fmt"

	"github.com/ShramikPG/CricketSimulationinGo/matchdetails"
	"github.com/jmcvetta/randutil"
	"github.com/logrusorgru/aurora"
)

// Service Interface to enforce commentory
type Service interface {
	Commentory() string
}

// Commentory returns the commentory as string
func Commentory(player matchdetails.Player, event string, score matchdetails.MatchScore, scoredRuns int) string {
	str := ""
	if event == "Out" {
		str = "\n Oh my God a fatal mistake by " + player.Name + " and he is out, What a blow. \n"
		return str
	}
	if event == "Scored" {
		switch scoredRuns {
		case 4:
			str = "\n What a wonderful boundary. " + player.Name + " is absolutely unstopabble, 4 runs for him \n"
		case 6:
			str = "\n What a hit son, what a hit. A maximum. " + player.Name + " you beauty, can he be stopped mates. \n"
		case 1:
			str = "\n Settling down for a single, keeping the wickets in hand, smart move by " + player.Name + "\n"
		case 2:
			str = "\n Cheeky, both the players stole two runs, and fielders couldn't do nothing about it." + player.Name + " gets two runs in his runBank and the team's\n"
		case 3:
			str = "\n Oh my word, these players have wheels in their legs, awesome running between the wicket. 3 Runs right off the fielders noses. " + player.Name + " Will gladly take these 3 runs in his scoresheet\n"
		case 5:
			str = "\n Five runs off the bat of " + player.Name + ". He will get the single, team will get the \n 5 as 4 is of the overthrow that leaded to the boundary"
		default:
			str = "\n Technincal details, sorry for the trouble \n"
		}
	}
	return str
}

//UpdateAfterDelivery updates Score after each ball
func UpdateAfterDelivery(event randutil.Choice, playerOnStrike, playerNotOnStrike *matchdetails.Player,
	matchScore *matchdetails.MatchScore) {
	switch event.Item {
	case 1:
		playerOnStrike.UpdateDetails(1, false)
		playerOnStrike.ChangeStrike()
		playerNotOnStrike.ChangeStrike()
		matchScore.UpdateDetails(1, 1, 0)
		fmt.Println(aurora.Cyan(Commentory(*playerOnStrike, "Scored", *matchScore, 1)))
	case 2:
		playerOnStrike.UpdateDetails(2, false)
		matchScore.UpdateDetails(2, 1, 0)
		fmt.Println(aurora.Magenta(Commentory(*playerOnStrike, "Scored", *matchScore, 2)))
	case 3:
		playerOnStrike.UpdateDetails(3, false)
		playerOnStrike.ChangeStrike()
		playerNotOnStrike.ChangeStrike()
		matchScore.UpdateDetails(3, 1, 0)
		fmt.Println(aurora.Yellow(Commentory(*playerOnStrike, "Scored", *matchScore, 3)))
	case 4:
		playerOnStrike.UpdateDetails(4, false)
		matchScore.UpdateDetails(4, 1, 0)
		fmt.Println(aurora.BrightGreen(Commentory(*playerOnStrike, "Scored", *matchScore, 4)))
	case 5:
		playerOnStrike.UpdateDetails(5, false)
		playerOnStrike.ChangeStrike()
		playerNotOnStrike.ChangeStrike()
		matchScore.UpdateDetails(5, 1, 0)
		fmt.Println(aurora.Blue(Commentory(*playerOnStrike, "Scored", *matchScore, 5)))
	case 6:
		playerOnStrike.UpdateDetails(6, false)
		matchScore.UpdateDetails(6, 1, 0)
		fmt.Println(aurora.Green(Commentory(*playerOnStrike, "Scored", *matchScore, 6)))
	case "Out":
		playerOnStrike.UpdateDetails(0, true)
		matchScore.UpdateDetails(0, 1, 1)
		fmt.Println(aurora.Red(Commentory(*playerOnStrike, "Out", *matchScore, 0)))
	}
}

// ChangeWickets is used to change the player after a wicket is taken
func ChangeWickets(currentPlayers [2]*matchdetails.Player, wicketNo9, wicketNo10 *matchdetails.Player,
	matchScore *matchdetails.MatchScore) {
	switch matchScore.WicketsLeft {
	case 2:
		if currentPlayers[0].IsPlaying == false {
			currentPlayers[0] = wicketNo9
			currentPlayers[0].IsPlaying = true
		} else {
			currentPlayers[1] = wicketNo9
			currentPlayers[1].IsPlaying = true
		}
	case 1:
		if currentPlayers[0].IsPlaying == false {
			currentPlayers[0] = wicketNo10
			currentPlayers[0].IsPlaying = true
		} else {
			currentPlayers[1] = wicketNo10
			currentPlayers[1].IsPlaying = true
		}
	case 0:
		matchScore.MatchVerdict()
	}
}
