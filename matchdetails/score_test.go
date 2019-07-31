package matchdetails

import (
	"testing"
)

func TestMatchScore_UpdateDetails(t *testing.T) {
	type fields struct {
		WicketsLeft    int
		RunsLeft       int
		DeliveriesLeft int
	}
	type args struct {
		decrementrunsby     int
		decrementDeliveryby int
		decrementWicketsBy  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"test1", fields{1, 4, 2}, args{1, 1, 0}},
		{"test2", fields{3, 40, 24}, args{0, 1, 1}},
		{"test3", fields{2, 34, 20}, args{6, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matchScore := &MatchScore{
				WicketsLeft:    tt.fields.WicketsLeft,
				RunsLeft:       tt.fields.RunsLeft,
				DeliveriesLeft: tt.fields.DeliveriesLeft,
			}
			matchScore.UpdateDetails(tt.args.decrementrunsby, tt.args.decrementDeliveryby, tt.args.decrementWicketsBy)
			if matchScore.RunsLeft != tt.fields.RunsLeft-tt.args.decrementrunsby {
				t.Errorf("Failed to Update Runs in MatchScore")
			}
			if matchScore.WicketsLeft != tt.fields.WicketsLeft-tt.args.decrementWicketsBy {
				t.Errorf("Failed to Update Wickets in MatchScore")
			}
			if matchScore.DeliveriesLeft != tt.fields.DeliveriesLeft-tt.args.decrementDeliveryby {
				t.Errorf("Failed to Update Delieveries in MatchScore")
			}
		})
	}
}

func TestMatchScore_OversLeft(t *testing.T) {
	type fields struct {
		WicketsLeft    int
		RunsLeft       int
		DeliveriesLeft int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  int
	}{
		{"Test1", fields{2, 40, 24}, 4, 0},
		{"Test2", fields{2, 40, 23}, 3, 5},
		{"Test1", fields{2, 40, 22}, 3, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matchScore := &MatchScore{
				WicketsLeft:    tt.fields.WicketsLeft,
				RunsLeft:       tt.fields.RunsLeft,
				DeliveriesLeft: tt.fields.DeliveriesLeft,
			}
			got, got1 := matchScore.OversLeft()
			if got != tt.want {
				t.Errorf("MatchScore.OversLeft() got = %v, want %v", got, tt.want)
			}

			if got1 != tt.want1 {
				t.Errorf("MatchScore.OversLeft() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMatchScore_MatchVerdict(t *testing.T) {
	type fields struct {
		WicketsLeft    int
		RunsLeft       int
		DeliveriesLeft int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matchScore := MatchScore{
				WicketsLeft:    tt.fields.WicketsLeft,
				RunsLeft:       tt.fields.RunsLeft,
				DeliveriesLeft: tt.fields.DeliveriesLeft,
			}
			if got := matchScore.MatchVerdict(); got != tt.want {
				t.Errorf("MatchScore.MatchVerdict() = %v, want %v", got, tt.want)
			}
		})
	}
}
