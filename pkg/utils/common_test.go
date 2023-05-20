package utils

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestParseUUID(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want uuid.UUID
	}{
		// TODO: Add test cases.
		{
			name: "happy flow: parse uuid from string",
			args: args{
				in: "a8a3b5e0-0b1a-4b1a-8f1a-0e1a0b1a0e1a",
			},
			want: uuid.MustParse("a8a3b5e0-0b1a-4b1a-8f1a-0e1a0b1a0e1a"),
		},
		{
			name: "bad flow: parse uuid from string",
			args: args{
				in: "8f1a-0e1a0b1a0e1a",
			},
			want: uuid.Nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseUUID(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
