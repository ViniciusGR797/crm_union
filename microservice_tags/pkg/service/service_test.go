package service

import (
	"microservice_tags/pkg/database"
	"reflect"
	"testing"
)

func TestNewTagsService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Tags_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTagsService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagsService() = %v, want %v", got, tt.want)
			}
		})
	}
}
