package http

import "fmt"

type Header map[string]interface{}

const Version string = "HTTP/1.1"

type Protocol interface {
	Message() string
}

type Request struct {
	Method  string
	Path    string
	Version string
}

func NewRequest(method, path, version string) *Request {
	return &Request{
		Method:  method,
		Path:    path,
		Version: version,
	}
}

func (m *Request) Message() string {
	return fmt.Sprintf("%s %s %s", m.Method, m.Path, m.Version)
}

type Response struct {
	Version    string
	StatusCode int
	StatusText string
}

func NewResponse(version string, statusCode int, statusText string) *Response {
	return &Response{
		Version:    version,
		StatusCode: statusCode,
		StatusText: statusText,
	}
}

func (m *Response) Message() string {
	return fmt.Sprintf("%s %d %s", m.Version, m.StatusCode, m.StatusText)
}

type Http struct {
	StartLine string
	Headers   []Header
	Body      []byte
}

func (h *Http) startLine(prt Protocol) {
	h.StartLine = prt.Message()
}

func (h *Http) header(key string, value interface{}) {
	h.Headers = append(h.Headers, Header{
		key: value,
	})
}

func HeaderStr(header Header) string {
	result := ""
	for key, value := range header {
		if _, ok := value.(string); ok {
			result += fmt.Sprintf("%s: %s\n", key, value)
		} else if _, ok = value.(int); ok {
			result += fmt.Sprintf("%s: %d\n", key, value)
		}
	}
	return result
}

func (h *Http) body(data string) {
	h.Body = []byte(data)
}

func (h *Http) Format() []byte {
	format := fmt.Sprintf("%s\n%s\n%s\n", h.StartLine, h.FormatHeaders(), h.Body)
	fmt.Println(format)
	return []byte(format)
}

func (h *Http) ContentType(contentType string) {
	h.header("Content-Type", contentType)
}

func (h *Http) ContentLength(contentLength int) {
	h.header("Content-Length", contentLength)
}

func (h *Http) headers(headers []Header) {
	for _, header := range headers {
		for k, v := range header {
			h.header(k, v)
		}
	}
}

func (h *Http) FormatHeaders() string {
	response := ""
	for _, header := range h.Headers {
		response += HeaderStr(header)
	}
	return response
}

func FormatRequest(message, path, body string, headers []Header) *Http {
	req := NewRequest(message, path, Version)
	return buildMessage(req, body, headers)
}

func FormatResponse(body string, statusCode int, statusText string, headers []Header) *Http {
	res := NewResponse(Version, statusCode, statusText)
	return buildMessage(res, body, headers)
}

func buildMessage(prt Protocol, body string, headers []Header) *Http {
	http := &Http{}
	http.startLine(prt)
	http.headers(headers)
	http.ContentLength(len(body))
	http.body(body)
	return http
}
