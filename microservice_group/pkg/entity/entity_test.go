package entity

import (
	"testing"
)

func TestGroup_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Group
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Group.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupID_String(t *testing.T) {
	tests := []struct {
		name string
		p    *GroupID
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("GroupID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustumer_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Custumer
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Custumer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_String(t *testing.T) {
	tests := []struct {
		name string
		p    *User
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("User.String() = %v, want %v", got, tt.want)
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

func TestCreateGroup_String(t *testing.T) {
	tests := []struct {
		name string
		p    *CreateGroup
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("CreateGroup.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountUser_String(t *testing.T) {
	tests := []struct {
		name string
		p    *CountUser
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("CountUser.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountUsersList_String(t *testing.T) {
	tests := []struct {
		name string
		p    *CountUsersList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("CountUsersList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_String(t *testing.T) {
	tests := []struct {
		name string
		p    *ID
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("ID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupIDList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *GroupIDList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("GroupIDList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *GroupList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("GroupList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *UserList
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("UserList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
