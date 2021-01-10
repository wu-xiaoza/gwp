package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlePost(t *testing.T) {
	//创建一个用于测试的多路复用器
	mux := http.NewServeMux()
	//绑定想要测试的处理器
	mux.HandleFunc("/post/", handleRequest)

	//创建记录器，用于获取服务器返回的HTTP响应
	writer := httptest.NewRecorder()
	//为被测试的处理器创建相应的请求
	json := strings.NewReader(`{"content":"Posted post","aauthor":"Da Hong"}`)

	request, _ := http.NewRequest("POST", "/post/1", json)
	//向被测试的处理器发送请求
	mux.ServeHTTP(writer, request)

	//对记录器记载的结果进行检查
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
