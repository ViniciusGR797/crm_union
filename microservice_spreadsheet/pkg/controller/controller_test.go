package controller

import (
	"microservice_spreadsheet/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGenerateSpreadSheet(t *testing.T) {
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
			GenerateSpreadSheet(tt.args.c, tt.args.service)
		})
	}
}
