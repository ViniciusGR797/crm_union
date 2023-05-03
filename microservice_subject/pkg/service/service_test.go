package service

import (
	"microservice_subject/pkg/database"
	"reflect"
	"testing"
)

func TestNewGroupService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Subject_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGroupService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroupService() = %v, want %v", got, tt.want)
			}
		})
	}
}
