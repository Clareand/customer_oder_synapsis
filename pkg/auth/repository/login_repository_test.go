package repository

import (
	"reflect"
	"testing"

	"github.com/Clareand/customer_oder_synapsis/pkg/auth/model"
)

func Test_loginRepo_CheckUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		r    *loginRepo
		args args
		want model.CheckUserIsTrue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.CheckUser(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginRepo.CheckUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginRepo_GetPassword(t *testing.T) {
	type args struct {
		userId string
	}
	tests := []struct {
		name         string
		r            *loginRepo
		args         args
		wantPassword string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPassword := tt.r.GetPassword(tt.args.userId); gotPassword != tt.wantPassword {
				t.Errorf("loginRepo.GetPassword() = %v, want %v", gotPassword, tt.wantPassword)
			}
		})
	}
}
