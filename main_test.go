package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHelloWorld(t *testing.T) {
	// リクエストを生成
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerHelloWorld)
	handler.ServeHTTP(recorder, req)

	// レスポンスを検証
	expected := "hello world"
	if recorder.Body.String() != expected {
		t.Errorf("Got: %v, Expected: %v", recorder.Body.String(), expected)
	}

	// ステータスコードの検証
	if recorder.Code != http.StatusOK {
		t.Errorf("Got status code: %v, Expected: %v", recorder.Code, http.StatusOK)
	}
}
