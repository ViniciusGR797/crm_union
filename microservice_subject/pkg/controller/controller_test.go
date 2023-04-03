package controller

import (
	"microservice_subject/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetSubmissiveSubjects(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.SubjectServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSubmissiveSubjects(tt.args.c, tt.args.service)
		})
	}
}

func TestGetSubjectByID(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.SubjectServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSubjectByID(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusSubjectFinished(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.SubjectServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusSubjectFinished(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateStatusSubjectCanceled(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.SubjectServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatusSubjectCanceled(tt.args.c, tt.args.service)
		})
	}
}

func TestCreateSubject(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.SubjectServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateSubject(tt.args.c, tt.args.service)
		})
	}
}

func TestUpdateSubject(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.SubjectServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateSubject(tt.args.c, tt.args.service)
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
