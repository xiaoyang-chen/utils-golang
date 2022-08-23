package utils

import (
	"testing"
	"time"
)

func TestNextTime(t *testing.T) {

	var spec, strWantNext = "0 0 21 * * *", "2022-08-23 21:00:00"

	var loc, _ = time.LoadLocation("Asia/Shanghai")
	var wantNext, _ = time.ParseInLocation("2006-01-02 15:04:05", strWantNext, loc)

	type args struct {
		spec  string
		start time.Time
	}
	tests := []struct {
		name     string
		args     args
		wantNext time.Time
		wantErr  bool
	}{
		{
			name: "test",
			args: args{
				spec:  spec,
				start: time.Now(),
			},
			wantNext: wantNext,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNext, err := NextTime(tt.args.spec, tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("NextTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotNext, tt.wantNext) {
			// 	t.Errorf("NextTime() = %v, want %v", gotNext, tt.wantNext)
			// }
			if !gotNext.Equal(tt.wantNext) {
				t.Errorf("NextTime() = %v, want %v", gotNext, tt.wantNext)
			}
		})
	}
}

func TestTickerRun(t *testing.T) {
	type args struct {
		spec  string
		start time.Time
		run   func(errTickerRun error, now time.Time)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				spec:  "0/2 * * * * *",
				start: time.Now(),
				run: func(errTickerRun error, now time.Time) {
					if errTickerRun != nil {
						t.Logf("errTickerRun: %v", errTickerRun)
						return
					}
					t.Log(now)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TickerRun(tt.args.spec, tt.args.start, tt.args.run)
			go TickerRun(tt.args.spec, tt.args.start, tt.args.run)
		})
	}
	<-time.After(10 * time.Second)
}
