package factory_mode

import "fmt"

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
