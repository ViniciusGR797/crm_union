package database

import (
	"database/sql"
	"microservice_subject/config"
	"reflect"
	"testing"
)

func TestNewDB(t *testing.T) {
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
			if got := NewDB(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dabase_pool_Close(t *testing.T) {
	tests := []struct {
		name    string
		d       *dabase_pool
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Close(); (err != nil) != tt.wantErr {
				t.Errorf("dabase_pool.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dabase_pool_GetDB(t *testing.T) {
	tests := []struct {
		name   string
		d      *dabase_pool
		wantDB *sql.DB
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDB := tt.d.GetDB(); !reflect.DeepEqual(gotDB, tt.wantDB) {
				t.Errorf("dabase_pool.GetDB() = %v, want %v", gotDB, tt.wantDB)
			}
		})
	}
}
