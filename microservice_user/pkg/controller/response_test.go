package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_sendError(t *testing.T) {
	type args struct {
		c      *gin.Context
		status int
		err    error
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendError(tt.args.c, tt.args.status, tt.args.err)
		})
	}
}

func Test_send(t *testing.T) {
	type args struct {
		c    *gin.Context
		code int
		obj  any
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			send(tt.args.c, tt.args.code, tt.args.obj)
		})
	}
}

func Test_sendNoContent(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendNoContent(tt.args.c)
		})
	}
}
