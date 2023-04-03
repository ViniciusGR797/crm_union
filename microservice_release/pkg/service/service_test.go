package service

import (
	"microservice_release/pkg/database"
	"microservice_release/pkg/entity"
	"reflect"
	"testing"
)

func TestNewReleaseService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Release_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReleaseService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReleaseService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelease_service_GetReleasesTrain(t *testing.T) {
	tests := []struct {
		name    string
		ps      *Release_service
		want    *entity.ReleaseList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetReleasesTrain()
			if (err != nil) != tt.wantErr {
				t.Errorf("Release_service.GetReleasesTrain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Release_service.GetReleasesTrain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelease_service_GetReleaseTrainByID(t *testing.T) {
	type args struct {
		ID uint64
	}
	tests := []struct {
		name    string
		ps      *Release_service
		args    args
		want    *entity.Release
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetReleaseTrainByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Release_service.GetReleaseTrainByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Release_service.GetReleaseTrainByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelease_service_UpdateReleaseTrain(t *testing.T) {
	type args struct {
		ID      uint64
		release *entity.Release_Update
		logID   *int
	}
	tests := []struct {
		name    string
		ps      *Release_service
		args    args
		want    uint64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.UpdateReleaseTrain(tt.args.ID, tt.args.release, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Release_service.UpdateReleaseTrain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Release_service.UpdateReleaseTrain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelease_service_GetTagsReleaseTrain(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *Release_service
		args    args
		want    []*entity.Tag
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetTagsReleaseTrain(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Release_service.GetTagsReleaseTrain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Release_service.GetTagsReleaseTrain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelease_service_InsertTagsReleaseTrain(t *testing.T) {
	type args struct {
		ID    uint64
		tags  []entity.Tag
		logID *int
	}
	tests := []struct {
		name    string
		ps      *Release_service
		args    args
		want    uint64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.InsertTagsReleaseTrain(tt.args.ID, tt.args.tags, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Release_service.InsertTagsReleaseTrain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Release_service.InsertTagsReleaseTrain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelease_service_UpdateStatusReleaseTrain(t *testing.T) {
	type args struct {
		ID    *uint64
		logID *int
	}
	tests := []struct {
		name    string
		ps      *Release_service
		args    args
		want    int64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.UpdateStatusReleaseTrain(tt.args.ID, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Release_service.UpdateStatusReleaseTrain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Release_service.UpdateStatusReleaseTrain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelease_service_GetReleaseTrainByBusiness(t *testing.T) {
	type args struct {
		businessID *uint64
	}
	tests := []struct {
		name    string
		ps      *Release_service
		args    args
		want    *entity.ReleaseList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetReleaseTrainByBusiness(tt.args.businessID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Release_service.GetReleaseTrainByBusiness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Release_service.GetReleaseTrainByBusiness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelease_service_CreateReleaseTrain(t *testing.T) {
	type args struct {
		release *entity.Release_Update
		logID   *int
	}
	tests := []struct {
		name    string
		ps      *Release_service
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ps.CreateReleaseTrain(tt.args.release, tt.args.logID); (err != nil) != tt.wantErr {
				t.Errorf("Release_service.CreateReleaseTrain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
