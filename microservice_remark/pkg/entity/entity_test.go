package entity

import (
	"testing"
)

func TestRemark_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Remark
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Remark.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemarkList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *RemarkList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("RemarkList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
