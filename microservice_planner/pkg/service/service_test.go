package service

import (
	"microservice_planner/pkg/database"
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
