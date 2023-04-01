package middlewares

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Auth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthAdmin(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthAdmin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}
