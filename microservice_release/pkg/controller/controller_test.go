package controller

import (
	"microservice_release/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetReleasesTrain(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ReleaseServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetReleasesTrain(tt.args.c, tt.args.service)
		})
	}
}

func TestGetReleaseTrainByID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ReleaseServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetReleaseTrainByID(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateReleaseTrain(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ReleaseServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateReleaseTrain(tt.args.c, tt.args.service)
		})
	}
}

func TestGetTagsReleaseTrain(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ReleaseServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTagsReleaseTrain(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusReleaseTrain(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ReleaseServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusReleaseTrain(tt.args.c, tt.args.service)
		})
	}
}

func TestGetReleaseTrainByBusiness(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ReleaseServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetReleaseTrainByBusiness(tt.args.c, tt.args.service)
		})
	}
}

func TestCreateReleaseTrain(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ReleaseServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateReleaseTrain(tt.args.c, tt.args.service)
		})
	}
}
