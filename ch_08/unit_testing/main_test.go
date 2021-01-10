package main

import (
	"testing"	
	"time"
)


func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}

	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}

	if post.Content != "Hello World!" {
		t.Error("Wrong content, was expecting 'Hello World!' but got", post.Content)
	}

	if post.Comments[0].Author != "Zhong Liang" {
		t.Error("Wrong comments[0].author, was expecting 'Zhong Liang' but got", post.Comments[0].Author)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")//Skip is equivalent to Log followed by SkipNow.//跳过测试
}

func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}