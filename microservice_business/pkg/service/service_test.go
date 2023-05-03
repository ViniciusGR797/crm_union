package service

import (
	"microservice_business/pkg/database"
	"reflect"
	"testing"
)

func TestNewBusinessService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Business_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBusinessService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBusinessService() = %v, want %v", got, tt.want)
			}
		})
	}
}
