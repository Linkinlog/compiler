package main

import "testing"

func Test_main(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
	}{
		{
			name: "Test_main_01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
