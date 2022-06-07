package builder

// 建造者模式之链式调用，用于简单的构建，每个属性的赋值逻辑比较简单。
// 通过链式调用和建造者模式结合，让调用者的可读性更好。

import "io"

type RequestOpts struct {
	MoreHeaders  map[string]interface{}
	JSONResponse interface{}
	JSONRaw      io.Reader
	OKCodes      []int
	KeepResponse bool
}

type RequestOptsBuilder RequestOpts

func NewRequestOptsBuilder() *RequestOptsBuilder {
	return &RequestOptsBuilder{}
}

func (b *RequestOptsBuilder) SetMoreHeaders(v map[string]interface{}) *RequestOptsBuilder {
	if v == nil {
		return b
	}
	b.MoreHeaders = v
	return b
}

func (b *RequestOptsBuilder) SetJSONResponse(v interface{}) *RequestOptsBuilder {
	if v == nil {
		return b
	}
	b.JSONResponse = v
	return b
}

func (b *RequestOptsBuilder) SetJSONRaw(v io.Reader) *RequestOptsBuilder {
	if v == nil {
		return b
	}
	b.JSONRaw = v
	return b
}

func (b *RequestOptsBuilder) SetOKCodes(v []int) *RequestOptsBuilder {
	if v == nil {
		return b
	}
	b.OKCodes = v
	return b
}

func (b *RequestOptsBuilder) SetKeepResponse(v bool) *RequestOptsBuilder {
	b.KeepResponse = v
	return b
}

func (b *RequestOptsBuilder) Build() *RequestOpts {
	// 在这里可以做一些设置默认值的操作
	if b.MoreHeaders == nil {
		b.MoreHeaders = make(map[string]interface{})
	}

	return (*RequestOpts)(b)
}
