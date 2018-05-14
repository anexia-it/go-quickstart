package quickstart

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Request represents a message queue request
type Request struct {
	// URL to fetch
	URL string `json:"url"`
	// Timeout in milliseconds
	TimeoutMillis uint `json:"timeout_ms"`
}

// ToString returns a string represenation of the request
func (r Request) ToString() string {
	timeoutDuration := time.Millisecond * time.Duration(r.TimeoutMillis)
	return fmt.Sprintf("url=%s, timeout=%s", r.URL, timeoutDuration)
}

// DecodeRequest decodes a request
// Returns an error and a nil pointer on error or a filled request struct on success
func DecodeRequest(r io.Reader) (req *Request, err error) {
	dec := json.NewDecoder(r)
	req = &Request{}
	if err = dec.Decode(req); err != nil {
		req = nil
	}
	return
}
