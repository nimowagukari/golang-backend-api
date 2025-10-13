package main

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h := &helloWorldHandler{}

	h.ServeHTTP(w, req)
	resp := w.Result()

	// ステータスコード
	if resp.StatusCode != 200 {
		t.Errorf("got: %d, want: %d", resp.StatusCode, 200)
	}

	// レスポンスボディ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(body, []byte("Hello,World.")) {
		t.Errorf("got: %s, want: %s", body, []byte("Hello,World."))
	}
}
