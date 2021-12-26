package utils

import (
	"testing"
)

func TestGetSqlParamFromStrings(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name         string
		args         args
		wantSqlParam string
	}{
		{
			name: "nil strings",
			args: args{
				strings: nil,
			},
			wantSqlParam: "",
		},
		{
			name: "empty strings",
			args: args{
				strings: []string{},
			},
			wantSqlParam: "",
		},
		{
			name: "len(strings) == 1",
			args: args{
				strings: []string{"123saf"},
			},
			wantSqlParam: "'123saf'",
		},
		{
			name: "len(strings) == 2",
			args: args{
				strings: []string{"123saf", "12312af12"},
			},
			wantSqlParam: "'123saf', '12312af12'",
		},
		{
			name: "len(strings) == 3",
			args: args{
				strings: []string{"123saf", "12312af12", "eda213"},
			},
			wantSqlParam: "'123saf', '12312af12', 'eda213'",
		},
		{
			name: "len(strings) == 4",
			args: args{
				strings: []string{"123saf", "12312af12", "eda213", "eda13z'c"},
			},
			wantSqlParam: "'123saf', '12312af12', 'eda213', 'eda13zc'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSqlParam := GetSqlParamFromStrings(tt.args.strings); gotSqlParam != tt.wantSqlParam {
				t.Errorf("GetSqlParamFromStrings() = %v, want %v", gotSqlParam, tt.wantSqlParam)
			}
		})
	}
}

func TestGetSqlParamFromIntegers(t *testing.T) {
	type args struct {
		integers interface{}
	}
	tests := []struct {
		name         string
		args         args
		wantSqlParam string
	}{
		{
			name: "integers: nil",
			args: args{
				integers: nil,
			},
			wantSqlParam: "",
		},
		{
			name: "integers: (*int)(nil)",
			args: args{
				integers: (*int)(nil),
			},
			wantSqlParam: "",
		},
		{
			name: "integers: []uint32 len 1",
			args: args{
				integers: []uint32{123},
			},
			wantSqlParam: "123",
		},
		{
			name: "integers: []uint32 len 2",
			args: args{
				integers: []uint32{123, 23},
			},
			wantSqlParam: "123,23",
		},
		{
			name: "integers: []uint32 len 3",
			args: args{
				integers: []uint32{123, 23, 22},
			},
			wantSqlParam: "123,23,22",
		},
		{
			name: "integers: []uint32 len 4",
			args: args{
				integers: []uint32{123, 23, 22, 33},
			},
			wantSqlParam: "123,23,22,33",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSqlParam := GetSqlParamFromIntegers(tt.args.integers); gotSqlParam != tt.wantSqlParam {
				t.Errorf("GetSqlParamFromIntegers() = %v, want %v", gotSqlParam, tt.wantSqlParam)
			}
		})
	}
}

func TestGetLikeSqlParamStrInTwoPercent(t *testing.T) {
	type args struct {
		likeStr string
	}
	tests := []struct {
		name         string
		args         args
		wantSqlParam string
	}{
		{
			name: "empty likeStr",
			args: args{
				likeStr: "",
			},
			wantSqlParam: "",
		},
		{
			name: "2dca2''1",
			args: args{
				likeStr: "2dca2''1",
			},
			wantSqlParam: "",
		},
		{
			name: "souls",
			args: args{
				likeStr: "souls",
			},
			wantSqlParam: "'%souls%'",
		},
		{
			name: "1fca21",
			args: args{
				likeStr: "1fca21",
			},
			wantSqlParam: "'%1fca21%'",
		},
		{
			name: "2dca21",
			args: args{
				likeStr: "2dca21",
			},
			wantSqlParam: "'%2dca21%'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSqlParam := GetLikeSqlParamStrInTwoPercent(tt.args.likeStr); gotSqlParam != tt.wantSqlParam {
				t.Errorf("GetLikeSqlParamStrInTwoPercent() = %v, want %v", gotSqlParam, tt.wantSqlParam)
			}
		})
	}
}

func TestGetSqlParamStrInTwoSingleQuote(t *testing.T) {
	type args struct {
		srcStr string
	}
	tests := []struct {
		name         string
		args         args
		wantSqlParam string
	}{
		{
			name: "empty srcStr",
			args: args{
				srcStr: "",
			},
			wantSqlParam: "''",
		},
		{
			name: "2dca2''1",
			args: args{
				srcStr: "2dca2''1",
			},
			wantSqlParam: "",
		},
		{
			name: "souls",
			args: args{
				srcStr: "souls",
			},
			wantSqlParam: "'souls'",
		},
		{
			name: "1fca21",
			args: args{
				srcStr: "1fca21",
			},
			wantSqlParam: "'1fca21'",
		},
		{
			name: "2dca21",
			args: args{
				srcStr: "2dca21",
			},
			wantSqlParam: "'2dca21'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSqlParam := GetSqlParamStrInTwoSingleQuote(tt.args.srcStr); gotSqlParam != tt.wantSqlParam {
				t.Errorf("GetSqlParamStrInTwoSingleQuote() = %v, want %v", gotSqlParam, tt.wantSqlParam)
			}
		})
	}
}

func TestGetMulQuestionMarkDividedByComma(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name         string
		args         args
		wantSqlParam string
	}{
		{
			name: "length < 0",
			args: args{
				length: -1,
			},
			wantSqlParam: "",
		},
		{
			name: "length == 0",
			args: args{
				length: 0,
			},
			wantSqlParam: "",
		},
		{
			name: "length == 1",
			args: args{
				length: 1,
			},
			wantSqlParam: "?",
		},
		{
			name: "length == 2",
			args: args{
				length: 2,
			},
			wantSqlParam: "?,?",
		},
		{
			name: "length == 3",
			args: args{
				length: 3,
			},
			wantSqlParam: "?,?,?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSqlParam := GetMulQuestionMarkDividedByComma(tt.args.length); gotSqlParam != tt.wantSqlParam {
				t.Errorf("SqlParam() = %v, want %v", gotSqlParam, tt.wantSqlParam)
			}
		})
	}
}
