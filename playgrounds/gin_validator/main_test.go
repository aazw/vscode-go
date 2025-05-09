// main_test.go
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"testing"
)

func performRequest(t *testing.T, handler http.Handler, method, apiPath, body string) *http.Response {

	// サーバを起動（HTTP/2 は切ると確実に HTTP/1.1 になる）
	ts := httptest.NewUnstartedServer(handler)
	ts.EnableHTTP2 = false
	ts.Start()
	defer ts.Close()

	httpClient := &http.Client{}
	u, _ := url.Parse(ts.URL)
	u.Path = path.Join(u.Path, apiPath)

	req, _ := http.NewRequest(method, u.String(), bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	// リクエストのダンプ
	if dumpReq, err := httputil.DumpRequest(req, true); err == nil {
		fmt.Fprint(os.Stdout, string(dumpReq)+"\n")
	} else {
		t.Logf("DumpRequestOut failed: %v\n", err)
	}
	fmt.Println()

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("http request failed: %v\n", err)
	}
	defer resp.Body.Close()

	// レスポンスのダンプ
	if dumpResp, err := httputil.DumpResponse(resp, true); err == nil {
		fmt.Fprint(os.Stdout, string(dumpResp)+"\n")
	} else {
		t.Fatalf("DumpResponse failed: %v\n", err)
	}

	return resp
}

func TestBindHandler_NoTest_OutputOnly(t *testing.T) {
	r := setupRouter()

	paths := make([]string, 0, len(handlerFuncs))
	for _, pair := range handlerFuncs {
		paths = append(paths, pair.Path)
	}
	body := `{}`

	for index, path := range paths {

		performRequest(t, r, "POST", path, body)

		if index+1 != len(paths) {
			fmt.Println("\n==========")
		}
	}
}
