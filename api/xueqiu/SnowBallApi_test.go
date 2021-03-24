package xueqiu

import (
	"testing"
)

func TestGetTopHandle(t *testing.T) {
	type args struct {
		symbol string
	}
	case1 := args{
		symbol: "sdf",
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"symbol", case1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTopHolder(tt.args.symbol)
		})
	}
}
