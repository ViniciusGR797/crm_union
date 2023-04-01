package controller

import (
	"microservice_group/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetGroups(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.GroupServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetGroups(tt.args.c, tt.args.service)
		})
	}
}

func TestGetGroupByID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.GroupServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetGroupByID(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusGroup(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.GroupServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusGroup(tt.args.c, tt.args.service)
		})
	}
}

func TestGetUsersGroup(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.GroupServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUsersGroup(tt.args.c, tt.args.service)
		})
	}
}

func TestCreateGroup(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.GroupServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateGroup(tt.args.c, tt.args.service)
		})
	}
}

func TestAttachUserGroup(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.GroupServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AttachUserGroup(tt.args.c, tt.args.service)
		})
	}
}

func TestDetachUserGroup(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.GroupServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DetachUserGroup(tt.args.c, tt.args.service)
		})
	}
}

func TestCountUsersGroup(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.GroupServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountUsersGroup(tt.args.c, tt.args.service)
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
