package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	// Place your code here
	dir := "./testdata/env"
	tests := map[string]EnvValue{
		"BAR": {
			Value:      "bar",
			NeedRemove: false,
		},
		"EMPTY": {
			Value:      "",
			NeedRemove: false,
		},
		"FOO": {
			Value:      "   foo\nwith new line",
			NeedRemove: false,
		},
		"HELLO": {
			Value:      "\"hello\"",
			NeedRemove: false,
		},
		"UNSET": {
			Value:      "",
			NeedRemove: false,
		},
	}
	got, err := ReadDir(dir)
	fmt.Println(got)
	if err != nil {
		return
	}
	for k := range tests {
		fmt.Println(tests[k], got[k])
		ok := reflect.DeepEqual(tests[k], got[k])
		require.Equal(t, ok, true)
	}
}
