package decodestream

import (
	"io"
	"reflect"
	"testing"
)

func TestNewJSONStream(t *testing.T) {
	tests := []struct {
		name string
		want Stream
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJSONStream(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream_Start(t *testing.T) {
	type fields struct {
		stream chan Entry
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stream{
				stream: tt.fields.stream,
			}
			s.Start(tt.args.r)
		})
	}
}

func TestStream_Watch(t *testing.T) {
	type fields struct {
		stream chan Entry
	}
	tests := []struct {
		name   string
		fields fields
		want   <-chan Entry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stream{
				stream: tt.fields.stream,
			}
			if got := s.Watch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Watch() = %v, want %v", got, tt.want)
			}
		})
	}
}
