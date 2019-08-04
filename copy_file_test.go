package main

import (
	"testing"
)

func TestCopy(t *testing.T) {
	type args struct {
		from   string
		to     string
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "File not found", args: args{from: "10.txt", to: "1.txt", offset: 480, limit: 3}, wantErr: true},
		{name: "Offset is bigger than size of file", args: args{from: "1.txt", to: "2.txt", offset: 4800, limit: 3}, wantErr: true},
		{name: "Offset + Limit is bigger than size of file", args: args{from: "1.txt", to: "3.txt", offset: 480, limit: 300}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.from, tt.args.to, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
