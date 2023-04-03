package entity

import (
	"testing"
)

func TestSubject_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Subject
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Subject.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubjectID_String(t *testing.T) {
	tests := []struct {
		name string
		p    *SubjectID
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("SubjectID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Client
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Client.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Status
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Status.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_list_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Subject_list
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Subject_list.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSubject_String(t *testing.T) {
	tests := []struct {
		name string
		p    *CreateSubject
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("CreateSubject.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomain_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Domain
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Domain.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
