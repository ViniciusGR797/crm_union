package service

import (
	"microservice_client/pkg/database"
	"microservice_client/pkg/entity"
	"reflect"
	"testing"
)

func TestNewClientService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Client_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_service_GetClientsMyGroups(t *testing.T) {
	type args struct {
		ID *int
	}
	tests := []struct {
		name    string
		ps      *Client_service
		args    args
		want    *entity.ClientList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetClientsMyGroups(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client_service.GetClientsMyGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client_service.GetClientsMyGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_service_GetClientByID(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *Client_service
		args    args
		want    *entity.Client
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetClientByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client_service.GetClientByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client_service.GetClientByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_service_GetClientByReleaseID(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *Client_service
		args    args
		want    *entity.ClientList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetClientByReleaseID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client_service.GetClientByReleaseID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client_service.GetClientByReleaseID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_service_GetTagsClient(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *Client_service
		args    args
		want    []*entity.Tag
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetTagsClient(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client_service.GetTagsClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client_service.GetTagsClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_service_GetRoles(t *testing.T) {
	tests := []struct {
		name string
		ps   *Client_service
		want *entity.RoleList
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetRoles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client_service.GetRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}
