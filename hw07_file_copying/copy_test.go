package main

import (
	"crypto/md5"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy1(t *testing.T) {
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "0_limit0", args: struct {
				fromPath string
				toPath   string
				offset   int64
				limit    int64
			}{fromPath: "./testdata/input.txt", toPath: "./out", offset: 0, limit: 0},
			wantErr: false,
		},
		{
			name: "0_limit10", args: struct {
				fromPath string
				toPath   string
				offset   int64
				limit    int64
			}{fromPath: "./testdata/input.txt", toPath: "./out", offset: 0, limit: 10},
			wantErr: false,
		},
		{
			name: "0_limit1000", args: struct {
				fromPath string
				toPath   string
				offset   int64
				limit    int64
			}{fromPath: "./testdata/input.txt", toPath: "./out", offset: 0, limit: 1000},
			wantErr: false,
		},
		{
			name: "0_limit10000", args: struct {
				fromPath string
				toPath   string
				offset   int64
				limit    int64
			}{fromPath: "./testdata/input.txt", toPath: "./out", offset: 0, limit: 10000},
			wantErr: false,
		},
		{
			name: "100_limit1000", args: struct {
				fromPath string
				toPath   string
				offset   int64
				limit    int64
			}{fromPath: "./testdata/input.txt", toPath: "./out", offset: 100, limit: 1000},
			wantErr: false,
		},
		{
			name: "6000_limit1000", args: struct {
				fromPath string
				toPath   string
				offset   int64
				limit    int64
			}{fromPath: "./testdata/input.txt", toPath: "./out", offset: 6000, limit: 10000},
			wantErr: false,
		},
		{
			name: "60000000_limit1000", args: struct {
				fromPath string
				toPath   string
				offset   int64
				limit    int64
			}{fromPath: "./testdata/input.txt", toPath: "./out", offset: 60000000, limit: 10000},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.fromPath, tt.args.toPath, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
				result, err := os.OpenFile(tt.args.toPath, os.O_RDONLY, 0o644)
				if err != nil {
					t.Error(err)
				}
				checkfile, err := os.OpenFile("./testdata/out_offset"+tt.name+".txt", os.O_RDONLY, 0o644)
				if err != nil {
					t.Error(err)
				}
				defer func() {
					result.Close()
					checkfile.Close()
					os.Remove(tt.args.toPath)
				}()

				hashDst := md5.New()
				hashSrt := md5.New()
				_, err = io.Copy(hashDst, checkfile)
				require.NoError(t, err)
				_, err = io.Copy(hashSrt, result)
				require.NoError(t, err)
				require.Equal(t, string(hashSrt.Sum(nil)), string(hashSrt.Sum(nil)))
			}
		})
	}
}
