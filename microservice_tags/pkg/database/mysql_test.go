package database

import (
	"microservice_tags/config"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMysql(t *testing.T) {
	type args struct {
		conf *config.Config
	}
	tests := []struct {
		name string
		args args
		want *dabase_pool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mysql(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mysql() = %v, want %v", got, tt.want)
			}
		})
	}
}
