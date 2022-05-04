package factory

// 当创建的对象不是简单的new一下就可以，而是需要做其他各种初始化操作的时候，
// 可以使用工厂方法模式，为每个复杂的对象创建一个相应的工厂，然后这个工厂会创建出实例，此实例

type RequestMessage struct {
	Body   []byte
	Helper RequestHelper
}

func (rm *RequestMessage) Read() string {
	rm.Helper.Help()
	return string(rm.Body)
}

type ResponseMessage struct {
	Body   []byte
	Helper ResponseHelper
}

func (rm *ResponseMessage) Read() string {
	rm.Helper.Help()
	return string(rm.Body)
}

// 创建每个对象的工厂。

type Factoryer interface {
	CreateMessage() Reader
}

type RequestMessageFactory struct{}

func (f RequestMessageFactory) CreateMessage() Reader {
	// 在这里可以进行各种初始化。
	body := RequestBody()
	helper := RequestHelper{}
	return &RequestMessage{
		Body:   body,
		Helper: helper,
	}
}

type ResponseMessageFactory struct{}

func (f ResponseMessageFactory) CreateMessage() Reader {
	// 在这里可以进行各种初始化操作
	body := ResponseBody()
	helper := ResponseHelper{}
	return &ResponseMessage{
		Body:   body,
		Helper: helper,
	}
}

type Factory struct{}

func NewFactoryer(mt string) Factoryer {
	switch mt {
	case "request":
		return RequestMessageFactory{}
	case "response":
		return ResponseMessageFactory{}
	}
	return nil
}
