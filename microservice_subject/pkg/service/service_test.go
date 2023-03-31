package service

import (
	"microservice_subject/pkg/database"
	"microservice_subject/pkg/entity"
	"reflect"
	"testing"
)

func TestNewGroupService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Subject_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGroupService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroupService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_service_GetSubmissiveSubjects(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		s       *Subject_service
		args    args
		want    *entity.Subject_list
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetSubmissiveSubjects(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subject_service.GetSubmissiveSubjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subject_service.GetSubmissiveSubjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_service_GetSubjectByID(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		s       *Subject_service
		args    args
		want    *entity.SubjectID
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetSubjectByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subject_service.GetSubjectByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subject_service.GetSubjectByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_service_UpdateStatusSubjectFinished(t *testing.T) {
	type args struct {
		id    uint64
		logID *int
	}
	tests := []struct {
		name    string
		s       *Subject_service
		args    args
		want    int64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateStatusSubjectFinished(tt.args.id, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subject_service.UpdateStatusSubjectFinished() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Subject_service.UpdateStatusSubjectFinished() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_service_UpdateStatusSubjectCanceled(t *testing.T) {
	type args struct {
		id    uint64
		logID *int
	}
	tests := []struct {
		name    string
		s       *Subject_service
		args    args
		want    int64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateStatusSubjectCanceled(tt.args.id, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subject_service.UpdateStatusSubjectCanceled() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Subject_service.UpdateStatusSubjectCanceled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_service_CreateSubject(t *testing.T) {
	type args struct {
		subject *entity.CreateSubject
		id      uint64
		logID   *int
	}
	tests := []struct {
		name    string
		s       *Subject_service
		args    args
		want    *entity.SubjectID
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateSubject(tt.args.subject, tt.args.id, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subject_service.CreateSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subject_service.CreateSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_service_UpdateSubject(t *testing.T) {
	type args struct {
		id      uint64
		subject *entity.UpdateSubject
		logID   *int
	}
	tests := []struct {
		name    string
		s       *Subject_service
		args    args
		want    int64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateSubject(tt.args.id, tt.args.subject, tt.args.logID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subject_service.UpdateSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Subject_service.UpdateSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}
