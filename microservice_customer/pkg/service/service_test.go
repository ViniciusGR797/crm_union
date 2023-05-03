package service

import (
	"microservice_customer/pkg/database"
	"reflect"
	"testing"
)

func TestNewCostumerService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *customer_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCostumerService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCostumerService() = %v, want %v", got, tt.want)
			}
		})
	}
}
