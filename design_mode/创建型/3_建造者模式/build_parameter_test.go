package builder

import (
	"reflect"
	"testing"
)

func TestEndpointOptionsBuilder(t *testing.T) {
	setService := EndpointOptionsFunc(func(option *EndpointOptions) {
		option.Service = "network"
	})
	setType := EndpointOptionsFunc(func(option *EndpointOptions) {
		option.Type = "internal"
	})
	setRegion := EndpointOptionsFunc(func(option *EndpointOptions) {
		option.Region = "CHINA_CHINA"
	})

	opt, err := NewEndpointOptions(setService, setType, setRegion)
	if err != nil {
		t.Fatal(err)
	}

	expect := EndpointOptions{
		Service: "network",
		Type:    "internal",
		Region:  "CHINA_CHINA",
	}

	if !reflect.DeepEqual(expect, *opt) {
		t.Fatalf("failed: want %v but got %v", expect, opt)
	}
}
