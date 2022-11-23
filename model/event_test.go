package model

import (
	"fmt"
	"testing"
)

func TestStrToTime(t *testing.T) {
	type args struct {
		timeStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_empty_str",
			args: args{
				timeStr: "",
			},
			wantErr: true,
		},
		{
			name: "test_year",
			args: args{
				timeStr: "3997-12-09 15:04",
			},
			wantErr: false,
		},
		{
			name: "test_month_out_of_range",
			args: args{
				timeStr: "3997-22-09",
			},
			wantErr: true,
		},
		{
			name: "test_day_out_of_range",
			args: args{
				timeStr: "3997-12-99 01:09",
			},
			wantErr: true,
		},
		{
			name: "test_hour_out_of_range",
			args: args{
				timeStr: "3997-12-99 25:00",
			},
			wantErr: true,
		},
		{
			name: "test_minute_out_of_range",
			args: args{
				timeStr: "3997-12-99 23:99",
			},
			wantErr: true,
		},
		{
			name: "test_hour_at_24",
			args: args{
				timeStr: "3997-12-09 24:07",
			},
			wantErr: true,
		},
		{
			name: "test_hour_minute_60",
			args: args{
				timeStr: "3997-12-09 23:60",
			},
			wantErr: true,
		},
		{
			name: "test_february_29_correct",
			args: args{
				timeStr: "2016-02-29 23:03",
			},
			wantErr: false,
		},
		{
			name: "test_february_29_wrong",
			args: args{
				timeStr: "2015-02-29 23:03",
			},
			wantErr: true,
		},
		{
			name: "test_letter",
			args: args{
				timeStr: "1997-12-0x 09:07",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StrToTime(tt.args.timeStr)
			fmt.Println(got, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
