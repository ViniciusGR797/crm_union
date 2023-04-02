package entity

import (
	"reflect"
	"testing"
)

func TestClient_String(t *testing.T) {
	tests := []struct {
		name string
		c    *Client
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Client.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *RoleList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("RoleList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientList_String(t *testing.T) {
	tests := []struct {
		name string
		cl   *ClientList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cl.String(); got != tt.want {
				t.Errorf("ClientList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		client_id    uint64
		client_name  string
		client_email string
		client_role  uint64
		customer_id  uint64
		release_id   uint64
	}
	tests := []struct {
		name string
		args args
		want *ClientUpdate
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.client_id, tt.args.client_name, tt.args.client_email, tt.args.client_role, tt.args.customer_id, tt.args.release_id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
