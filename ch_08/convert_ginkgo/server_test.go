package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("get post", func() {

		mux := http.NewServeMux()
		mux.HandleFunc("/post/", handleRequest(&FakePost{}))
		writer := httptest.NewRecorder()

		request, _ := http.NewRequest("GET", "/post/1", nil)
		mux.ServeHTTP(writer, request)

		if writer.Code != 200 {
			GinkgoT().Errorf("Response code is %v", writer.Code)
		}
		var post Post
		json.Unmarshal(writer.Body.Bytes(), &post)
		if post.Id != 1 {
			GinkgoT().Errorf("Cannot retrieve JSON post")
		}
	})
	It("put post", func() {

		mux := http.NewServeMux()
		post := &FakePost{}
		mux.HandleFunc("/post/", handleRequest(post))

		writer := httptest.NewRecorder()
		json := strings.NewReader(`{"content":"Updated post","author":"Da Hong"}`)
		request, _ := http.NewRequest("PUT", "/post/1", json)
		mux.ServeHTTP(writer, request)

		if writer.Code != 200 {
			GinkgoT().Errorf("Response code is %v", writer.Code)
		}

		if post.Content != "Updated post" {
			GinkgoT().Errorf("Content is not correct", post.Content)
		}
	})
})
