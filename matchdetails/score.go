package matchdetails

import "strconv"

// MatchScore struct saves the current state of match
type MatchScore struct {
	WicketsLeft    int
	RunsLeft       int
	DeliveriesLeft int
}

const (
	// NotOver just to tell main that match is not over yet
	NotOver = "Not over yet"
)

// UpdateDetails is used to update match details
func (matchScore *MatchScore) UpdateDetails(decrementrunsby int, decrementDeliveryby int,
	decrementWicketsBy int) {
	if decrementrunsby > 0 {
		matchScore.RunsLeft -= decrementrunsby
	}
	if decrementWicketsBy > 0 {
		matchScore.WicketsLeft -= decrementWicketsBy
	}
	if decrementDeliveryby > 0 {
		matchScore.DeliveriesLeft -= decrementDeliveryby
	}
}

// OversLeft returns overs left and balls left in that over respectively
func (matchScore *MatchScore) OversLeft() (int, int) {
	if matchScore.DeliveriesLeft < 0 {
		return 0, 0
	}
	oversLeft := matchScore.DeliveriesLeft / 6
	delRemThisOver := matchScore.DeliveriesLeft % 6
	return oversLeft, delRemThisOver
}

// MatchVerdict gives the  result of the match
func (matchScore MatchScore) MatchVerdict() string {
	res := ""
	if matchScore.WicketsLeft <= 0 || matchScore.DeliveriesLeft <= 0 {
		res = "\nLengaburu Lost the match by " + strconv.Itoa(matchScore.RunsLeft) + " Runs. \n"
		return res
	}
	if matchScore.RunsLeft <= 0 {
		res = "\nLengaburu won by " + strconv.Itoa(matchScore.WicketsLeft) + " wickets and " + strconv.Itoa(matchScore.DeliveriesLeft) + " balls in hand. \n"
		return res
	}
	return NotOver
}
