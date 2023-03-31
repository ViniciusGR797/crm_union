package security

import "testing"

func TestRandStringRunes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandStringRunes(tt.args.n); got != tt.want {
				t.Errorf("RandStringRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
