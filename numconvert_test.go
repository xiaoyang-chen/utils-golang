package utils

import (
	"math"
	"testing"
)

func TestNumToFloat64(t *testing.T) {
	type args struct {
		a interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 bool
	}{
		{
			name: "10.123",
			args: args{
				a: 10.123,
			},
			want:  10.123,
			want1: true,
		},
		{
			name: "10",
			args: args{
				a: 10,
			},
			want:  10,
			want1: true,
		},
		{
			name: "string",
			args: args{
				a: "10",
			},
			want:  0,
			want1: false,
		},
		{
			name: "slice",
			args: args{
				a: []int{1, 2},
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NumToFloat64(tt.args.a)
			if got != tt.want {
				t.Errorf("NumToFloat64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NumToFloat64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFloat64ToKMGT(t *testing.T) {
	type args struct {
		float float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "< 1k",
			args: args{
				float: 9.8763,
			},
			want: "9.8763",
		},
		{
			name: "> -1k && < 0",
			args: args{
				float: -9.8763,
			},
			want: "-9.8763",
		},
		{
			name: "< 1k",
			args: args{
				float: 9.87634356546,
			},
			want: "9.87634356546",
		},
		{
			name: "> -1k && < 0",
			args: args{
				float: -9.87634356546,
			},
			want: "-9.87634356546",
		},
		{
			name: "< 1k",
			args: args{
				float: 9.00,
			},
			want: "9",
		},
		{
			name: "> -1k && < 0",
			args: args{
				float: -9.00,
			},
			want: "-9",
		},
		{
			name: "< 1k",
			args: args{
				float: 9.0900,
			},
			want: "9.09",
		},
		{
			name: "> -1k && < 0",
			args: args{
				float: -9.0900,
			},
			want: "-9.09",
		},
		{
			name: "= 1k",
			args: args{
				float: 1000.00000,
			},
			want: "1.00K",
		},
		{
			name: "= -1k",
			args: args{
				float: -1000.00000,
			},
			want: "-1.00K",
		},
		{
			name: "1k < float < 1M",
			args: args{
				float: 1081.34200,
			},
			want: "1.08K",
		},
		{
			name: "> -1M && < -1k",
			args: args{
				float: -1081.34200,
			},
			want: "-1.08K",
		},
		{
			name: "> -1G && < -1M",
			args: args{
				float: -1081000.34200,
			},
			want: "-1.08M",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64ToKMGT(tt.args.float); got != tt.want {
				t.Errorf("Float64ToKMGT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConvertPositiveIntegerToStringBase10(t *testing.T) {
	type args struct {
		in int64
	}
	tests := []struct {
		name    string
		args    args
		wantStr string
	}{
		{
			name: "math.MaxInt64",
			args: args{
				in: math.MaxInt64,
			},
			wantStr: "9223372036854775807",
		},
		{
			name: "-111",
			args: args{
				in: -111,
			},
			wantStr: "-1",
		},
		{
			name: "1234567",
			args: args{
				in: 1234567,
			},
			wantStr: "1234567",
		},
		{
			name: "0",
			args: args{
				in: 0,
			},
			wantStr: "0",
		},
		{
			name: "1",
			args: args{
				in: 1,
			},
			wantStr: "1",
		},
		{
			name: "2",
			args: args{
				in: 2,
			},
			wantStr: "2",
		},
		{
			name: "3",
			args: args{
				in: 3,
			},
			wantStr: "3",
		},
		{
			name: "4",
			args: args{
				in: 4,
			},
			wantStr: "4",
		},
		{
			name: "5",
			args: args{
				in: 5,
			},
			wantStr: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStr := ConvertPositiveIntegerToStringBase10(tt.args.in); gotStr != tt.wantStr {
				t.Errorf("ConvertPositiveIntegerToStringBase10() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}
