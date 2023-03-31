package controller

import (
	"microservice_user/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUsers(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUsers(tt.args.c, tt.args.service)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUserByID(tt.args.c, tt.args.service)
		})
	}
}

func TestGetUserByName(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUserByName(tt.args.c, tt.args.service)
		})
	}
}

func TestGetSubmissiveUsers(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSubmissiveUsers(tt.args.c, tt.args.service)
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateUser(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusUser(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusUser(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateUser(tt.args.c, tt.args.service)
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Login(tt.args.c, tt.args.service)
		})
	}
}

func TestGetUserMe(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUserMe(tt.args.c, tt.args.service)
		})
	}
}
