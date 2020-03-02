package main

import (
	"reflect"
	"testing"

	"github.com/go-pg/pg"
)

func Test_errorf(t *testing.T) {
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "print error",
			args: args{
				s: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorf(tt.args.s, tt.args.args...)
		})
	}
}

func Test_getDBConnection(t *testing.T) {

	pgOpt := &pg.Options{
		Addr:     "localhost:5432",
		Database: "tokopedia-salamexp-quran",
		User:     "postgres",
		Password: "",
	}

	tests := []struct {
		name string
		want *pg.Options
	}{
		{
			name: "get pg configuration",
			want: pgOpt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDBConnection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDBConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
