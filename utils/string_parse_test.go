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

func TestCheckCNPhonesFormat(t *testing.T) {
	type args struct {
		phones string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "12345678901",
			args: args{
				phones: "12345678901",
			},
			wantErr: false,
		},
		{
			name: "12345678901,12345678901",
			args: args{
				phones: "12345678901,12345678901",
			},
			wantErr: false,
		},
		{
			name: "12345678901, 12345678901",
			args: args{
				phones: "12345678901, 12345678901",
			},
			wantErr: true,
		},
		{
			name: ",12345678901, 12345678901",
			args: args{
				phones: ",12345678901, 12345678901",
			},
			wantErr: true,
		},
		{
			name: "12345678901,1234567890",
			args: args{
				phones: "12345678901,1234567890",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckCNPhonesFormat(tt.args.phones); (err != nil) != tt.wantErr {
				t.Errorf("CheckCNPhonesFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
