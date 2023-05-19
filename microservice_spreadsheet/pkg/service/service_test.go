package service

import (
	"microservice_spreadsheet/pkg/database"
	"reflect"
	"testing"
)

func TestNewRemarkService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *remark_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRemarkService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRemarkService() = %v, want %v", got, tt.want)
			}
		})
	}
}
