package entity

import (
	"testing"
)

func TestRelease_String(t *testing.T) {
	tests := []struct {
		name string
		c    *Release
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Release.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReleaseList_String(t *testing.T) {
	tests := []struct {
		name string
		cl   *ReleaseList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cl.String(); got != tt.want {
				t.Errorf("ReleaseList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
