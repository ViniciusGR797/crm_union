package service

import (
	"microservice_tags/pkg/database"
	"microservice_tags/pkg/entity"
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

func TestTags_service_GetTags(t *testing.T) {
	tests := []struct {
		name string
		ps   *Tags_service
		want *entity.TagsList
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags_service.GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTags_service_GetTagsById(t *testing.T) {
	type args struct {
		ID uint64
	}
	tests := []struct {
		name    string
		ps      *Tags_service
		args    args
		want    *entity.Tags
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetTagsById(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tags_service.GetTagsById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags_service.GetTagsById() = %v, want %v", got, tt.want)
			}
		})
	}
}
