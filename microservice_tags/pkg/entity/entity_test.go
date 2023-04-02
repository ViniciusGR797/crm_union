package entity

import (
	"reflect"
	"testing"
)

func TestTags_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Tags
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Tags.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagsList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *TagsList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("TagsList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTag(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Tags
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTag(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
