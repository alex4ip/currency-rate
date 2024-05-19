package exchangerate

import (
	"net/http"
	"testing"
)

func TestRate(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Rate",
			args: args{
				w: nil,
				r: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rate(tt.args.w, tt.args.r)
		})
	}
}
