package sql

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestReadDBstring(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Totally not here for code coverage", args: args{filename: "testdb/test.sql"}, want: "Testerino", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDBstring(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDBstring() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadDBstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
