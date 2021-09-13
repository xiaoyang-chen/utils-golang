package utils

import (
	"math/rand"
	"testing"
)

func TestTimestampToLocalTimeStr(t *testing.T) {
	type args struct {
		ts int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "timestamp wrong",
			args: args{
				ts: rand.Int63n(1e9),
			},
			want: "timestamp wrong",
		},
		{
			name: "s级时间戳",
			args: args{
				ts: 1606798630,
			},
			want: "2020-12-01 12:57:10",
		},
		{
			name: "ms级时间戳",
			args: args{
				ts: 1606798630931,
			},
			want: "2020-12-01 12:57:10.931",
		},
		{
			name: "us级时间戳",
			args: args{
				ts: 1606798630931772,
			},
			want: "2020-12-01 12:57:10.931772",
		},
		{
			name: "ns级时间戳",
			args: args{
				ts: 1606798630931772168,
			},
			want: "2020-12-01 12:57:10.931772168",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimestampToLocalTimeStr(tt.args.ts); got != tt.want {
				t.Errorf("TimestampToLocalTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
