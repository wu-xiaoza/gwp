package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestHandleGet(t *testing.T) {
	//创建一个用于测试的多路复用器
	mux := http.NewServeMux()
	//绑定想要测试的处理器
	mux.HandleFunc("/post/", handleRequest)

	//创建记录器，用于获取服务器返回的HTTP响应
	writer := httptest.NewRecorder()
	//为被测试的处理器创建相应的请求
	request, _ := http.NewRequest("GET", "/post/1", nil)
	//向被测试的处理器发送请求
	mux.ServeHTTP(writer, request)

	//对记录器记载的结果进行检查
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"Da Hong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}