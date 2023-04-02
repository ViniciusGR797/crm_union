package entity

import (
	"reflect"
	"testing"
)

func TestUser_String(t *testing.T) {
	tests := []struct {
		name string
		p    *User
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("User.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *UserList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("UserList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		name       string
		email      string
		created_at string
		status     string
		level      uint
		id         uint64
		tcs_id     uint64
	}
	tests := []struct {
		name string
		args args
		want *User
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.name, tt.args.email, tt.args.created_at, tt.args.status, tt.args.level, tt.args.id, tt.args.tcs_id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Prepare(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.Prepare(); (err != nil) != tt.wantErr {
				t.Errorf("User.Prepare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_format(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.format(); (err != nil) != tt.wantErr {
				t.Errorf("User.format() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_validate(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.validate(); (err != nil) != tt.wantErr {
				t.Errorf("User.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
