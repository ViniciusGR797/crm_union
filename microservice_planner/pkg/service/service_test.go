package service

import (
	"microservice_planner/pkg/database"
	"microservice_planner/pkg/entity"
	"reflect"
	"testing"
)

func TestNewPlannerService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Planner_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlannerService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlannerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanner_service_GetPlannerByID(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *Planner_service
		args    args
		want    *entity.Planner
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetPlannerByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Planner_service.GetPlannerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Planner_service.GetPlannerByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanner_service_GetPlannerByName(t *testing.T) {
	type args struct {
		ID    *int
		level int
		name  *string
	}
	tests := []struct {
		name    string
		ps      *Planner_service
		args    args
		want    *entity.PlannerList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetPlannerByName(tt.args.ID, tt.args.level, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Planner_service.GetPlannerByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Planner_service.GetPlannerByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanner_service_GetSubmissivePlanners(t *testing.T) {
	type args struct {
		ID    *int
		level int
	}
	tests := []struct {
		name    string
		ps      *Planner_service
		args    args
		want    *entity.PlannerList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetSubmissivePlanners(tt.args.ID, tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("Planner_service.GetSubmissivePlanners() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Planner_service.GetSubmissivePlanners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanner_service_GetPlannerByBusiness(t *testing.T) {
	type args struct {
		name *string
	}
	tests := []struct {
		name    string
		ps      *Planner_service
		args    args
		want    *entity.PlannerList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetPlannerByBusiness(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Planner_service.GetPlannerByBusiness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Planner_service.GetPlannerByBusiness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanner_service_GetGuestClientPlanners(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *Planner_service
		args    args
		want    []*entity.Client
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetGuestClientPlanners(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Planner_service.GetGuestClientPlanners() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Planner_service.GetGuestClientPlanners() = %v, want %v", got, tt.want)
			}
		})
	}
}
