package decodestream

import (
	"encoding/json"
	"fmt"
	"io"
)

// Entry represent a stream. If the stream fails, an error will be present.
type Entry struct {
	Error error
	Data  interface{}
}

// Stream helps transmit each streams withing a channel.
type Stream struct {
	stream chan Entry
}

// NewJSONStream return a new `Stream` type.
func NewJSONStream() Stream {
	return Stream{
		stream: make(chan Entry),
	}
}

// Watch watches JSON streams. Each stream entry will either have an error or a port object
func (s Stream) Watch() <-chan Entry {
	return s.stream
}

// Start starts streaming JSON reader. If an error occurs, the stream channel
// will be closed.
func (s Stream) Start(r io.Reader) {

	// close the stream channel
	defer close(s.stream)

	// Decode JSON reader
	decoder := json.NewDecoder(r)
	i := 1
	for decoder.More() {
		//get the delimiter
		t, err := decoder.Token()
		if err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode delimiter: %w", err)}
			return
		}
		//if is not a delimiter, then decode the data
		if fmt.Sprintf("%T", t) != "json.Delim" {
			//decode the data
			var data interface{}
			err := decoder.Decode(&data)
			if err != nil {
				s.stream <- Entry{Error: fmt.Errorf("decode line %d: %w", i, err)}
				return
			}
			//send the data to the channel
			s.stream <- Entry{Data: data}
			i++
		}
	}
}
