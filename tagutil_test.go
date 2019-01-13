package tagutil

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		tag  string
		want Tag
	}{{
		tag: ``, want: nil,
	}, {
		tag: `json:"fieldName"`, want: Tag{"json": {"fieldName"}},
	}, {
		tag: `json:"-"`, want: Tag{"json": {"-"}},
	}, {
		tag:  `json:"field,omitempty" doc:"required"`,
		want: Tag{"json": {"field", "omitempty"}, "doc": {"required"}},
	}, {
		tag:  `json:",inline,omitempty"`,
		want: Tag{"json": {"", "inline", "omitempty"}},
	}}

	for _, tt := range tests {
		got := New(tt.tag)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
