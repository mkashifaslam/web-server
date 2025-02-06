package http

import "fmt"

type Header map[string]interface{}

const Version string = "HTTP/1.1"

type Message interface {
	Message() string
}

type RequestMessage struct {
	Method  string
	Path    string
	Version string
}

func NewRequestMessage(method, path, version string) *RequestMessage {
	return &RequestMessage{
		Method:  method,
		Path:    path,
		Version: version,
	}
}

func (m *RequestMessage) Message() string {
	return fmt.Sprintf("%s %s %s", m.Method, m.Path, m.Version)
}

type ResponseMessage struct {
	Version    string
	StatusCode int
	StatusText string
}

func NewResponseMessage(version string, statusCode int, statusText string) *ResponseMessage {
	return &ResponseMessage{
		Version:    version,
		StatusCode: statusCode,
		StatusText: statusText,
	}
}

func (m *ResponseMessage) Message() string {
	return fmt.Sprintf("%s %d %s", m.Version, m.StatusCode, m.StatusText)
}

type Http struct {
	StartLine string
	Headers   []Header
	Body      []byte
}

func (h *Http) startLine(message Message) {
	h.StartLine = message.Message()
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

func NewRequest(message, path, body string, headers []Header) *Http {
	req := NewRequestMessage(message, path, Version)
	http := &Http{}
	http.startLine(req)
	http.headers(headers)
	http.ContentLength(len(body))
	http.body(body)
	return http
}

func NewResponse(body string, statusCode int, statusText string, headers []Header) *Http {
	res := NewResponseMessage(Version, statusCode, statusText)
	http := &Http{}
	http.startLine(res)
	http.headers(headers)
	http.ContentLength(len(body))
	http.body(body)
	return http
}
