package builder

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 通过链式调用的方式适合简单赋值逻辑的对象使用，在复杂的赋值情况下可以使用下面的这种方式
// 比如需要校验参数等操作

type HttpReqeustBuilder http.Request

func (b *HttpReqeustBuilder) SetMethod(v string) error {
	validMethods := []string{http.MethodGet, http.MethodPost,
		http.MethodPatch, http.MethodDelete}

	for _, vv := range validMethods {
		if v == vv {
			b.Method = v
			return nil
		}
	}

	return fmt.Errorf("invalid http method")
}

func (b *HttpReqeustBuilder) SetURL(v *url.URL) error {
	if v.Scheme == "" {
		return fmt.Errorf("invalid request url: empty scheme")
	}
	if v.Host == "" {
		return fmt.Errorf("invalid request url: empty host")
	}
	b.URL = v
	return nil
}

func (b *HttpReqeustBuilder) SetHeader(v http.Header) error {
	// 在这里可以检查头部的一些信息
	if v == nil {
		return fmt.Errorf("the http header is nil")
	}
	b.Header = v
	return nil
}

func (b *HttpReqeustBuilder) SetBody(v io.Reader) error {
	if v == nil {
		return fmt.Errorf("the http body is nil")
	}
	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(v); err != nil {
		return fmt.Errorf("set http body failed: %v", err)
	}

	b.Body = io.NopCloser(buf)
	return nil
}

func (b *HttpReqeustBuilder) Build() *http.Request {
	// 在这里可以做一些设置默认值操作

	return (*http.Request)(b)
}
