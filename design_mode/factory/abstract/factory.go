package factory

import "fmt"

type Reader interface {
	Read() string
}

type Booker struct{}

func (b Booker) Read() string {
	return "read book"
}

type Writer interface {
	Write(string)
}

type Paper struct{}

func (p Paper) Write(v string) {
	fmt.Println("value is:", v)
}

type ReaderWriterFactoryer interface {
	CreateReader() Reader
	CreateWriter() Writer
}

type ReaderWriterFactory struct{}

func (f ReaderWriterFactory) CreateReader() Reader {
	return Booker{}
}

func (f ReaderWriterFactory) CreateWriter() Writer {
	return Paper{}
}
