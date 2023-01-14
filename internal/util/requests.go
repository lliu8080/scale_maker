package util

// See similar https://github.com/kirinlabs/HttpRequest/

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	cli               *http.Client
	transport         *http.Transport
	debug             bool
	url               string
	method            string
	time              int64
	timeout           time.Duration
	headers           map[string]string
	cookies           map[string]string
	data              interface{}
	disableKeepAlives bool
	tlsClientConfig   *tls.Config
	proxy             string
	checkRedirect     func(req *http.Request, via []*http.Request) error
}

func (r *Request) log() {
	if r.debug {
		fmt.Printf("\nRequest: %s %s\nHeaders: %v\nCookies: %v\nTimeout: %ds\nReqBody: %v\n\n", r.method, r.url, r.headers, r.cookies, r.timeout, r.data)
	}
}

func (r *Request) getDuration(startTime time.Time) int64 {
	duration := time.Since(startTime).Milliseconds()
	return duration
}

func (r *Request) Get(url string, data ...interface{}) (*Response, error) {
	return r.request(http.MethodGet, url, data...)
}

func (r *Request) Post(url string, data ...interface{}) (*Response, error) {
	return r.request(http.MethodPost, url, data...)
}

func (r *Request) Put(url string, data ...interface{}) (*Response, error) {
	return r.request(http.MethodPut, url, data...)
}

func (r *Request) Delete(url string, data ...interface{}) (*Response, error) {
	return r.request(http.MethodDelete, url, data...)
}

// Set headers
func (r *Request) SetHeaders(headers map[string]string) *Request {
	if headers != nil || len(headers) > 0 {
		for k, v := range headers {
			r.headers[k] = v
		}
	}
	return r
}

func (r *Request) getTransport() http.RoundTripper {
	if r.transport == nil {
		return http.DefaultTransport
	}

	r.transport.DisableKeepAlives = r.disableKeepAlives

	if r.tlsClientConfig != nil {
		r.transport.TLSClientConfig = r.tlsClientConfig
	}

	// if r.proxy != nil {
	// 	r.transport.Proxy = r.proxy
	// }

	return http.RoundTripper(r.transport)
}

func (r *Request) buildClient() *http.Client {
	if r.cli == nil {
		r.cli = &http.Client{
			Transport:     r.getTransport(),
			CheckRedirect: r.checkRedirect,
			Timeout:       time.Second * r.timeout,
		}
	}
	return r.cli
}

func (r *Request) request(method, url string, data ...interface{}) (*Response, error) {

	response := &Response{}

	start := time.Now()

	if method == "" || url == "" {
		return nil, errors.New("method and url are required!")
	}

	// Debug infomation
	defer r.log()

	r.url = url
	if len(data) > 0 {
		r.data = data[0]
	} else {
		r.data = ""
	}

	var (
		err  error
		req  *http.Request
		body io.Reader
	)
	r.cli = r.buildClient()

	method = strings.ToUpper(method)
	r.method = method

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := r.cli.Do(req)
	if err != nil {
		return nil, err
	}

	response.resp = resp
	response.duration = r.getDuration(start)
	response.url = url
	response.headers = resp.Header
	response.statusCode = resp.StatusCode
	response.body, _ = ioutil.ReadAll(resp.Body)

	return response, nil
}

func NewRequest() *Request {
	r := &Request{
		timeout: 30,
		headers: map[string]string{},
		cookies: map[string]string{},
	}
	return r
}

func HttpGetRequest(url string, timeOut int) (*http.Response, int64) {
	var httpTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(timeOut) * time.Second,
		}).Dial,
		TLSHandshakeTimeout: time.Duration(timeOut) * time.Second,
	}

	var httpClient = &http.Client{
		Timeout:   time.Duration(timeOut) * time.Second,
		Transport: httpTransport,
	}

	start := time.Now()
	response, _ := httpClient.Get(url)
	duration := time.Since(start).Milliseconds()

	return response, duration

}
