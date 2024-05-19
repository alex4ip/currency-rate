package exchangerate

import (
	"reflect"
	"testing"
)

func TestGetRates(t *testing.T) {
	type args struct {
		valcode string
		date    string
	}
	tests := []struct {
		name    string
		args    args
		want    []ExchangeRate
		wantErr bool
	}{
		{
			name: "USD rate for date",
			args: args{
				"USD",
				"20200302",
			},
			want: []ExchangeRate{{R030: 840, Txt: "Долар США", Rate: 24.59, Cc: "USD", ExchangeDate: "02.03.2020"}},
		},
		{
			name: "USD rate today",
			args: args{
				"USD",
				"",
			},
			want: []ExchangeRate{{R030: 840, Txt: "Долар США", Rate: 39.4272, Cc: "USD", ExchangeDate: "20.05.2024"}},
		},

		{
			name: "EUR rate for date",
			args: args{
				"EUR",
				"20200302",
			},
			want: []ExchangeRate{{R030: 978, Txt: "Євро", Rate: 26.9789, Cc: "EUR", ExchangeDate: "02.03.2020"}},
		},
		{
			name: "EUR rate today",
			args: args{
				"EUR",
				"",
			},
			want: []ExchangeRate{{R030: 978, Txt: "Євро", Rate: 42.7549, Cc: "EUR", ExchangeDate: "20.05.2024"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRates(tt.args.valcode, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRate(t *testing.T) {
	type args struct {
		valcode string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "USD rate today float64",
			args:    args{"USD"},
			want:    39.4272,
			wantErr: false,
		},
		{
			name:    "000 rate 0 float64",
			args:    args{"000"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRate(tt.args.valcode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
