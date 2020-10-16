package main

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "positive number",
			args: args{
				x: 1,
			},
			wantResult: 2,
		},
		{
			name: "negative number",
			args: args{
				x: -2,
			},
			wantResult: -1,
		},
		{
			name: "zero",
			args: args{
				x: 0,
			},
			wantResult: 1,
		},
		{
			name: "one",
			args: args{
				x: 1,
			},
			wantResult: 2,
		},
		{
			name: "max int",
			args: args{
				x: 9223372036854775807,
			},
			wantResult: -9223372036854775808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Increment(tt.args.x); gotResult != tt.wantResult {
				t.Errorf("(%s) Increment() = %v, want %v", tt.name, gotResult, tt.wantResult)
			}
		})
	}
}
