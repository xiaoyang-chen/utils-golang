package utils

import (
	"math/big"
	"reflect"
	"testing"
)

func TestBigInt2Bytes(t *testing.T) {
	type args struct {
		b *big.Int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []byte
	}{
		{
			name: "nil",
			args: args{
				b: nil,
			},
			wantRes: nil,
		},
		{
			name: "0",
			args: args{
				b: BigIntZero(),
			},
			wantRes: []byte{},
		},
		{
			name: "-1",
			args: args{
				b: NewBigInt(-1),
			},
			wantRes: []byte{'-', 1},
		},
		{
			name: "1",
			args: args{
				b: NewBigInt(1),
			},
			wantRes: []byte{'+', 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := BigInt2Bytes(tt.args.b); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("BigInt2Bytes() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
