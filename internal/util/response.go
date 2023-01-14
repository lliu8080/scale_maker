package util

// See similar https://github.com/kirinlabs/HttpRequest/

import (
	"net/http"
)

type Response struct {
	resp       *http.Response
	duration   int64
	url        string
	headers    http.Header
	statusCode int
	body       []byte
}

func (r *Response) StatusCode() int {
	if r.resp == nil {
		return 0
	}
	return r.statusCode
}

func (r *Response) DurationMS() int64 {
	if r != nil {
		return r.duration
	}
	return 0
}

func (r *Response) Url() string {
	if r != nil {
		return r.url
	}
	return ""
}

func (r *Response) Headers() http.Header {
	if r.resp != nil {
		return r.resp.Header
	}
	return nil
}

func (r *Response) Content() string {
	if r.resp != nil {
		return string(r.body)
	}
	return ""
}

func (r *Response) Body() []byte {
	if r.resp != nil {
		return r.body
	}
	return nil
}
