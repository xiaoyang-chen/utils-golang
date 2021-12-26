package utils

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	strSample := []string{
		"你",
		"好",
		"世",
		"界",
		"语",
		"言",
		"变",
		"化",
		"hello",
		"world",
		"language",
		"change",
	}
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "abb",
			args: args{
				s: "abb",
			},
			want: []byte{'a', 'b', 'b'},
		},
	}
	for i := 0; i < 100; i++ {
		name := ""
		for j := 0; j < i+10; j++ {
			name += strSample[rand.Intn(len(strSample))]
		}
		tests = append(tests, struct {
			name string
			args args
			want []byte
		}{
			name: name,
			args: args{
				s: name,
			},
			want: ([]byte)(name),
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2Bytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str2Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes2Str(t *testing.T) {
	strSample := []string{
		"你",
		"好",
		"世",
		"界",
		"语",
		"言",
		"变",
		"化",
		"hello",
		"world",
		"language",
		"change",
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abb",
			args: args{
				b: []byte{'a', 'b', 'b'},
			},
			want: "abb",
		},
	}
	for i := 0; i < 100; i++ {
		name := ""
		for j := 0; j < i+10; j++ {
			name += strSample[rand.Intn(len(strSample))]
		}
		tests = append(tests, struct {
			name string
			args args
			want string
		}{
			name: name,
			args: args{
				b: ([]byte)(name),
			},
			want: (name),
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2Str(tt.args.b); got != tt.want {
				t.Errorf("Bytes2Str() = %v, want %v", got, tt.want)
			}
		})
	}
}
