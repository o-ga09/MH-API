package main

import "testing"

func Test_main(t *testing.T) {
	t.Skipf("main function is no test code . ")
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}