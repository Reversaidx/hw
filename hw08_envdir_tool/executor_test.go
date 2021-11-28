package main

import (
	"testing"
)

func TestRunCmd(t *testing.T) {
	type args struct {
		cmdS []string
		env  Environment
	}
	tests := []struct {
		name           string
		args           args
		wantReturnCode int
		wantStdout     string
	}{
		{
			name: "echo1",
			args: args{
				cmdS: []string{"/bin/bash", "-c", "echo", "-e", "${WORK_DIR}"},
				env: Environment{
					"TEST": EnvValue{
						Value:      "123",
						NeedRemove: false,
					},
				},
			},
			wantReturnCode: 0,
		},
		{
			name: "echo1",
			args: args{
				cmdS: []string{"echo", "$TEST"},
				env: Environment{
					"TEST": EnvValue{
						Value:      "123",
						NeedRemove: true,
					},
				},
			},
			wantReturnCode: 0,
		},
		{
			name: "dummpy",
			args: args{
				cmdS: []string{"/bin/bash", "-c", "./test2.sh", "$TEST"},
				env: Environment{
					"TEST": EnvValue{
						Value:      "123",
						NeedRemove: true,
					},
				},
			},
			wantReturnCode: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReturnCode := RunCmd(tt.args.cmdS, tt.args.env); gotReturnCode != tt.wantReturnCode {
				t.Errorf("RunCmd() = %v, want %v", gotReturnCode, tt.wantReturnCode)
			}
		})
	}
}
