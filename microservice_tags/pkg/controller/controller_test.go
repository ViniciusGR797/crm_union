package controller

import (
	"microservice_tags/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetTags(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.TagsServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTags(tt.args.c, tt.args.service)
		})
	}
}

func TestGetTagsById(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.TagsServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTagsById(tt.args.c, tt.args.service)
		})
	}
}
