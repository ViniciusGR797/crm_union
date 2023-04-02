package controller

import (
	"microservice_remark/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetSubmissiveRemarks(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.RemarkServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSubmissiveRemarks(tt.args.c, tt.args.service)
		})
	}
}

func TestGetAllRemarkUser(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.RemarkServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAllRemarkUser(tt.args.c, tt.args.service)
		})
	}
}

func TestGetRemarkByID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.RemarkServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRemarkByID(tt.args.c, tt.args.service)
		})
	}
}

func TestCreateRemark(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.RemarkServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateRemark(tt.args.c, tt.args.service)
		})
	}
}

func TestGetBarChartRemark(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.RemarkServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetBarChartRemark(tt.args.c, tt.args.service)
		})
	}
}

func TestGetPieChartRemark(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.RemarkServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPieChartRemark(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusRemark(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.RemarkServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusRemark(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateRemark(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.RemarkServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateRemark(tt.args.c, tt.args.service)
		})
	}
}

func TestJSONMessenger(t *testing.T) {
	type args struct {
		c      *gin.Context
		status int
		path   string
		err    error
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JSONMessenger(tt.args.c, tt.args.status, tt.args.path, tt.args.err)
		})
	}
}
