package controller

import (
	"microservice_planner/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetPlannerByID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.PlannerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPlannerByID(tt.args.c, tt.args.service)
		})
	}
}

func TestCreatePlanner(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.PlannerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePlanner(tt.args.c, tt.args.service)
		})
	}
}

func TestGetPlannerByName(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.PlannerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPlannerByName(tt.args.c, tt.args.service)
		})
	}
}

func TestGetSubmissivePlanners(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.PlannerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSubmissivePlanners(tt.args.c, tt.args.service)
		})
	}
}

func TestGetPlannerByBusiness(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.PlannerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPlannerByBusiness(tt.args.c, tt.args.service)
		})
	}
}

func TestGetGuestClientPlanners(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.PlannerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetGuestClientPlanners(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdatePlanner(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.PlannerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdatePlanner(tt.args.c, tt.args.service)
		})
	}
}
