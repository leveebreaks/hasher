// Package hasher provides functions to work with password hashing. SHA256 algorithm is being used.

package hasher

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"HashPassword Positive",
			args{password: "testinput"},
			"4LdZ8zau/S/1sxU08j2Yzt/cpAeFC6XGyZUCxCREGrc=",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"CheckPasswordHash equals",
			args{"testinput", "4LdZ8zau/S/1sxU08j2Yzt/cpAeFC6XGyZUCxCREGrc="},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
