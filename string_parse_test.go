package utils

import "testing"

func TestIsOnlyHasDigitalAndComma(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name          string
		args          args
		wantIsOnlyHas bool
	}{
		{
			name: "empty string",
			args: args{
				str: "",
			},
			wantIsOnlyHas: true,
		},
		{
			name: "123,,,123",
			args: args{
				str: "123,,,123",
			},
			wantIsOnlyHas: true,
		},
		{
			name: ",,,123",
			args: args{
				str: ",,,123",
			},
			wantIsOnlyHas: true,
		},
		{
			name: ",,a,123",
			args: args{
				str: ",,a,123",
			},
			wantIsOnlyHas: false,
		},
		{
			name: ",,.,123",
			args: args{
				str: ",,.,123",
			},
			wantIsOnlyHas: false,
		},
		{
			name: ",,，,123",
			args: args{
				str: ",,，,123",
			},
			wantIsOnlyHas: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsOnlyHas := IsOnlyHasDigitalAndComma(tt.args.str); gotIsOnlyHas != tt.wantIsOnlyHas {
				t.Errorf("IsOnlyHasDigitalAndComma() = %v, want %v", gotIsOnlyHas, tt.wantIsOnlyHas)
			}
		})
	}
}
