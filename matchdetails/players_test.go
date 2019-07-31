package matchdetails

import (
	"testing"

	"github.com/jmcvetta/randutil"
)

func TestPlayerUpdateDetails(t *testing.T) {
	type fields struct {
		Name           string
		CurrentScore   int
		IsPlaying      bool
		OnStrike       bool
		BallsPlayed    int
		ShotProbablity []randutil.Choice
	}
	type args struct {
		incrementscoreby int
		isOut            bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestUpdateDetails1",
			fields{"Kirat",
				10,
				true,
				true,
				5,
				[]randutil.Choice{{30, 0}, {25, 1}, {5, 2}, {0, 3}, {5, 4}, {1, 5}, {4, 6}, {30, "Out"}},
			},
			args{4, false},
		},
		{"TestUpdateDetails2",
			fields{"Kirat",
				10,
				true,
				true,
				5,
				[]randutil.Choice{{30, 0}, {25, 1}, {5, 2}, {0, 3}, {5, 4}, {1, 5}, {4, 6}, {30, "Out"}},
			},
			args{0, true},
		},
		{"TestUpdateDetails3",
			fields{"Kirat",
				10,
				true,
				true,
				5,
				[]randutil.Choice{{30, 0}, {25, 1}, {5, 2}, {0, 3}, {5, 4}, {1, 5}, {4, 6}, {30, "Out"}},
			},
			args{3, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := &Player{
				Name:           tt.fields.Name,
				CurrentScore:   tt.fields.CurrentScore,
				IsPlaying:      tt.fields.IsPlaying,
				OnStrike:       tt.fields.OnStrike,
				BallsPlayed:    tt.fields.BallsPlayed,
				ShotProbablity: tt.fields.ShotProbablity,
			}
			player.UpdateDetails(tt.args.incrementscoreby, tt.args.isOut)
			expectedScore := tt.fields.CurrentScore + tt.args.incrementscoreby
			if player.CurrentScore != expectedScore {
				t.Errorf("Unable to update score, expected %v, got %v", expectedScore, player.CurrentScore)
			}
			var expectedIsPlaying bool
			if tt.args.isOut == true {
				expectedIsPlaying = false
			} else {
				expectedIsPlaying = true
			}
			if player.IsPlaying != expectedIsPlaying {
				t.Errorf("Unable to update IsPlaying field, expected %v, got %v", expectedIsPlaying, player.IsPlaying)
			}
		})
	}
}

func TestPlayer_ChangeStrike(t *testing.T) {
	type fields struct {
		Name           string
		CurrentScore   int
		IsPlaying      bool
		OnStrike       bool
		BallsPlayed    int
		ShotProbablity []randutil.Choice
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestChangeStrike1",
			fields{"Kirat",
				10,
				true,
				true,
				5,
				[]randutil.Choice{{30, 0}, {25, 1}, {5, 2}, {0, 3}, {5, 4}, {1, 5}, {4, 6}, {30, "Out"}},
			},
		},
		{"TestChangeStrike1",
			fields{"Kirat",
				10,
				true,
				false,
				5,
				[]randutil.Choice{{30, 0}, {25, 1}, {5, 2}, {0, 3}, {5, 4}, {1, 5}, {4, 6}, {30, "Out"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := &Player{
				Name:           tt.fields.Name,
				CurrentScore:   tt.fields.CurrentScore,
				IsPlaying:      tt.fields.IsPlaying,
				OnStrike:       tt.fields.OnStrike,
				BallsPlayed:    tt.fields.BallsPlayed,
				ShotProbablity: tt.fields.ShotProbablity,
			}
			player.ChangeStrike()
			if player.OnStrike == tt.fields.OnStrike {
				t.Errorf("Error Updating Strike in ChangeStrike func expected %v, got %v", !tt.fields.OnStrike, player.OnStrike)
			}
		})
	}
}

func TestScoreCard(t *testing.T) {
	type args struct {
		players []Player
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ScoreCard(tt.args.players)
		})
	}
}
