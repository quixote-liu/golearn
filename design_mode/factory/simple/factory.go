package fatory

// 一般而言我们采用NewName来创建对象/接口，当它返回的是接口的时候，其实就是简单工厂模式。

type Booker interface {
	Name() string
	Type() string
}

type HistoryBook struct{}

func (b HistoryBook) Name() string {
	return "history book"
}

func (b HistoryBook) Type() string {
	return "history"
}

type PhysicsBook struct{}

func (b PhysicsBook) Name() string {
	return "physics book"
}

func (b PhysicsBook) Type() string {
	return "physics"
}

func NewBooker(book string) Booker {
	if book == "history" {
		return HistoryBook{}
	}
	if book == "physics" {
		return HistoryBook{}
	}

	return nil
}
