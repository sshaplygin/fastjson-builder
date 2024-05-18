package main

import (
	"encoding/json"
	"testing"

	"github.com/valyala/fastjson"
)

func BenchmarkStd(b *testing.B) {
	user := &User{ID: "test_value", Tags: []string{"1", "3", "2"}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&user)
	}
}

func BenchmarkManualFastJSONBuild(b *testing.B) {
	user := &User{ID: "test_value", Tags: []string{"1", "3", "2"}}

	pool := &fastjson.ArenaPool{}
	arena := pool.Get()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetUserVal(arena, user)
	}
}

func BenchmarkReflectFastJSONBuild(b *testing.B) {
	user := &User{ID: "test_value", Tags: []string{"1", "3", "2"}}

	pool := &fastjson.ArenaPool{}
	arena := pool.Get()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetReflectUserVal(arena, user)
	}
}
