package service

import (
	"microservice_business/pkg/database"
	"microservice_business/pkg/entity"
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

func TestBusiness_service_GetBusiness(t *testing.T) {
	tests := []struct {
		name string
		ps   *Business_service
		want *entity.BusinessList
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetBusiness(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Business_service.GetBusiness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBusiness_service_GetBusinessById(t *testing.T) {
	type args struct {
		ID uint64
	}
	tests := []struct {
		name    string
		ps      *Business_service
		args    args
		want    *entity.Business
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetBusinessById(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Business_service.GetBusinessById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Business_service.GetBusinessById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBusiness_service_GetBusinessByName(t *testing.T) {
	type args struct {
		name *string
	}
	tests := []struct {
		name    string
		ps      *Business_service
		args    args
		want    *entity.BusinessList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetBusinessByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Business_service.GetBusinessByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Business_service.GetBusinessByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBusiness_service_GetTagsBusiness(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *Business_service
		args    args
		want    []*entity.Tag
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetTagsBusiness(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Business_service.GetTagsBusiness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Business_service.GetTagsBusiness() = %v, want %v", got, tt.want)
			}
		})
	}
}
