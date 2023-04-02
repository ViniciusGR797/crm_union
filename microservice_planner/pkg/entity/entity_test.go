package entity

import (
	"testing"
)

func TestPlanner_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Planner
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Planner.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlannerList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *PlannerList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("PlannerList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanner_Prepare(t *testing.T) {
	tests := []struct {
		name    string
		p       *Planner
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Prepare(); (err != nil) != tt.wantErr {
				t.Errorf("Planner.Prepare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlanner_format(t *testing.T) {
	tests := []struct {
		name    string
		p       *Planner
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.format(); (err != nil) != tt.wantErr {
				t.Errorf("Planner.format() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlanner_validate(t *testing.T) {
	tests := []struct {
		name    string
		p       *Planner
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Planner.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
