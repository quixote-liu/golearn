package factory

import "fmt"

type Reader interface {
	Read() string
}

type Screen struct{}

func (s Screen) Read() string {
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
	return Screen{}
}

func (f ReaderWriterFactory) CreateWriter() Writer {
	return Paper{}
}

// ---------------- provide message ------------------

func RequestBody() []byte {
	return []byte("request body...")
}

func ResponseBody() []byte {
	return []byte("response body...")
}

type RequestHelper struct{}

func (rh RequestHelper) Help() {
	fmt.Println("help doing somethings...")
}

type ResponseHelper struct{}

func (rh ResponseHelper) Help() {
	fmt.Println("help doing somethings...")
}
