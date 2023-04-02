package controller

import (
	"microservice_customer/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetCustomers(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.CustomerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCustomers(tt.args.c, tt.args.service)
		})
	}
}

func TestGetCustomerByID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.CustomerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCustomerByID(tt.args.c, tt.args.service)
		})
	}
}

func TestCreateCustomer(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.CustomerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateCustomer(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.CustomerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCustomer(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusCustomer(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.CustomerServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusCustomer(tt.args.c, tt.args.service)
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
