package controller

import (
	"microservice_client/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetClientsMyGroups(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ClientServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetClientsMyGroups(tt.args.c, tt.args.service)
		})
	}
}

func TestGetClientByID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ClientServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetClientByID(tt.args.c, tt.args.service)
		})
	}
}

func TestGetClientByReleaseID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ClientServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetClientByReleaseID(tt.args.c, tt.args.service)
		})
	}
}

func TestGetTagsClient(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ClientServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTagsClient(tt.args.c, tt.args.service)
		})
	}
}

func TestCreateClient(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ClientServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateClient(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateClient(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ClientServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateClient(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusClient(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ClientServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusClient(tt.args.c, tt.args.service)
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

func TestGetRoles(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ClientServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRoles(tt.args.c, tt.args.service)
		})
	}
}
