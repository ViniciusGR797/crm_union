package service

import (
	"microservice_remark/pkg/database"
	"microservice_remark/pkg/entity"
	"reflect"
	"testing"
)

func TestNewRemarkService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *remark_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRemarkService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRemarkService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remark_service_GetSubmissiveRemarks(t *testing.T) {
	type args struct {
		ID *int
	}
	tests := []struct {
		name    string
		ps      *remark_service
		args    args
		want    *entity.RemarkList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetSubmissiveRemarks(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("remark_service.GetSubmissiveRemarks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remark_service.GetSubmissiveRemarks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remark_service_GetAllRemarkUser(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *remark_service
		args    args
		want    *entity.RemarkList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetAllRemarkUser(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("remark_service.GetAllRemarkUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remark_service.GetAllRemarkUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remark_service_GetRemarkByID(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *remark_service
		args    args
		want    *entity.Remark
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetRemarkByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("remark_service.GetRemarkByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remark_service.GetRemarkByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remark_service_GetBarChartRemark(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name string
		ps   *remark_service
		args args
		want *entity.Remark
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetBarChartRemark(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remark_service.GetBarChartRemark() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remark_service_GetPieChartRemark(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name string
		ps   *remark_service
		args args
		want *entity.Remark
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetPieChartRemark(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remark_service.GetPieChartRemark() = %v, want %v", got, tt.want)
			}
		})
	}
}
