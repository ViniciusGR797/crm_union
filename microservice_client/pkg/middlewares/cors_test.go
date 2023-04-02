package middlewares

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCORS(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CORS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CORS() = %v, want %v", got, tt.want)
			}
		})
	}
}
