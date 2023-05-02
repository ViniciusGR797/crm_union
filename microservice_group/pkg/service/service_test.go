package service

import (
	"microservice_group/pkg/database"
	"microservice_group/pkg/entity"
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
		want *Group_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGroupService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroupService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_service_GetGroups(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		ps      *Group_service
		args    args
		want    *entity.GroupList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetGroups(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Group_service.GetGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group_service.GetGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_service_GetGroupByID(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		ps      *Group_service
		args    args
		want    *entity.GroupID
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetGroupByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Group_service.GetGroupByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group_service.GetGroupByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_service_GetUsersGroup(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		ps      *Group_service
		args    args
		want    *entity.UserList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetUsersGroup(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Group_service.GetUsersGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group_service.GetUsersGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_service_CountUsersGroup(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		ps      *Group_service
		args    args
		want    *entity.CountUsersList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.CountUsersGroup(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Group_service.CountUsersGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group_service.CountUsersGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
