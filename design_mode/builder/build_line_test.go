package builder

import (
	"reflect"
	"testing"
)

func TestNewRequestOptsBuilder(t *testing.T) {
	okcodes := []int{200}
	jsonResponse := map[string]interface{}{}
	headers := map[string]interface{}{}

	builder := NewRequestOptsBuilder()
	opts := builder.SetOKCodes(okcodes).SetJSONResponse(jsonResponse).SetMoreHeaders(headers).Build()

	expect := RequestOpts{
		OKCodes:      okcodes,
		JSONResponse: jsonResponse,
		MoreHeaders:  headers,
	}

	if !reflect.DeepEqual(expect, *opts) {
		t.Fatalf("want %v but got %v", expect, opts)
	}
}
