package entity

import (
	"reflect"
	"testing"
)

func TestCustomer_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Customer
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Customer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *CustomerList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("CustomerList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCustomer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Customer
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomer(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
