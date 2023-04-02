package entity

import (
	"reflect"
	"testing"
)

func TestBusiness_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Business
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Business.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBusinessList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *BusinessList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("BusinessList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBusiness(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Business
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBusiness(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBusiness() = %v, want %v", got, tt.want)
			}
		})
	}
}
