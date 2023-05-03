package service

import (
	"microservice_release/pkg/database"
	"reflect"
	"testing"
)

func TestNewReleaseService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Release_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReleaseService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReleaseService() = %v, want %v", got, tt.want)
			}
		})
	}
}
