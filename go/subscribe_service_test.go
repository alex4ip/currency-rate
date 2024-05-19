package exchangerate

import (
	"database/sql"
	"testing"
)

func Test_createRatesTable(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "createRatesTable",
			args: args{
				db: nil,
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createRatesTable(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("createRatesTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createUsersTable(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createUsersTable(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("createUsersTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dbInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbInit()
		})
	}
}

func Test_subscribe(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subscribe(tt.args.email)
		})
	}
}
