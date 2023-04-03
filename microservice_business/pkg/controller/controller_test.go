package controller

import (
	"microservice_business/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetBusiness(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.BusinessServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetBusiness(tt.args.c, tt.args.service)
		})
	}
}

func TestGetBusinessById(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.BusinessServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetBusinessById(tt.args.c, tt.args.service)
		})
	}
}

func TestCreateBusiness(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.BusinessServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBusiness(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateBusiness(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.BusinessServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateBusiness(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusBusiness(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.BusinessServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusBusiness(tt.args.c, tt.args.service)
		})
	}
}

func TestGetBusinessByName(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.BusinessServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetBusinessByName(tt.args.c, tt.args.service)
		})
	}
}

func TestGetTagsBusiness(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.BusinessServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTagsBusiness(tt.args.c, tt.args.service)
		})
	}
}
