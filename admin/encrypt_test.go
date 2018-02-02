package main

import (
	"reflect"
	"testing"
)

func Test_generatePassword(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePassword(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
