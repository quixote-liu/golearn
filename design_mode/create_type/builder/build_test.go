package builder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHttpReqeustBuilder(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		message := map[string]interface{}{
			"method": r.Method,
			"header": r.Header,
			"body":   string(body),
		}

		jsonBytes, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}))

	builder := HttpReqeustBuilder{}

	u, _ := url.Parse(s.URL)
	if err := builder.SetURL(u); err != nil {
		t.Fatalf("build failed: url")
	}

	if err := builder.SetMethod(http.MethodPost); err != nil {
		t.Fatalf("build failed: method")
	}

	header := http.Header{}
	header.Set("hello", "server")
	if err := builder.SetHeader(header); err != nil {
		t.Fatalf("build failed: header")
	}

	body := bytes.NewBufferString("hello, builder")
	if err := builder.SetBody(body); err != nil {
		t.Fatalf("build failed: body")
	}

	req := builder.Build()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed")
	}

	actual := map[string]interface{}{}
	_ = json.NewDecoder(resp.Body).Decode(&actual)

	// 这里并没有进行断言，可以简单查看一下其中的信息
	// fmt.Println("actual:", actual)
}
