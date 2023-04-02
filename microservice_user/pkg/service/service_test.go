package service

import (
	"microservice_user/pkg/database"
	"microservice_user/pkg/entity"
	"reflect"
	"testing"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *User_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_service_GetUsers(t *testing.T) {
	tests := []struct {
		name    string
		ps      *User_service
		want    *entity.UserList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("User_service.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User_service.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_service_GetUserByID(t *testing.T) {
	type args struct {
		ID *int
	}
	tests := []struct {
		name    string
		ps      *User_service
		args    args
		want    *entity.User
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetUserByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("User_service.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User_service.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_service_GetUserByName(t *testing.T) {
	type args struct {
		name *string
	}
	tests := []struct {
		name    string
		ps      *User_service
		args    args
		want    *entity.UserList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetUserByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("User_service.GetUserByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User_service.GetUserByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_service_GetSubmissiveUsers(t *testing.T) {
	type args struct {
		ID    *int
		level int
	}
	tests := []struct {
		name    string
		ps      *User_service
		args    args
		want    *entity.UserList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetSubmissiveUsers(tt.args.ID, tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("User_service.GetSubmissiveUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User_service.GetSubmissiveUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_service_CreateUser(t *testing.T) {
	type args struct {
		user  *entity.User
		logID *int
	}
	tests := []struct {
		name    string
		ps      *User_service
		args    args
		want    uint64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.CreateUser(tt.args.user, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("User_service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User_service.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_service_UpdateStatusUser(t *testing.T) {
	type args struct {
		ID    *uint64
		logID *int
	}
	tests := []struct {
		name    string
		ps      *User_service
		args    args
		want    int64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.UpdateStatusUser(tt.args.ID, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("User_service.UpdateStatusUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User_service.UpdateStatusUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_service_UpdateUser(t *testing.T) {
	type args struct {
		ID    *int
		user  *entity.User
		logID *int
	}
	tests := []struct {
		name    string
		ps      *User_service
		args    args
		want    int
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.UpdateUser(tt.args.ID, tt.args.user, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("User_service.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User_service.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_service_Login(t *testing.T) {
	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		ps      *User_service
		args    args
		want    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.Login(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("User_service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User_service.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
