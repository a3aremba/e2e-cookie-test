package main

import "testing"

func TestInitMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitMsg(); got != tt.want {
				t.Errorf("InitMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
