package main

import (
	"reflect"
	"testing"
)

func TestReadDir(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name    string
		args    args
		want    Environment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDir(tt.args.directory)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunCmd(t *testing.T) {
	type args struct {
		cmdS []string
		env  Environment
	}
	tests := []struct {
		name           string
		args           args
		wantReturnCode int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReturnCode := RunCmd(tt.args.cmdS, tt.args.env); gotReturnCode != tt.wantReturnCode {
				t.Errorf("RunCmd() = %v, want %v", gotReturnCode, tt.wantReturnCode)
			}
		})
	}
}

func Test_parseEnv(t *testing.T) {
	type args struct {
		env []string
	}
	tests := []struct {
		name    string
		args    args
		want    Environment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEnv(tt.args.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}
