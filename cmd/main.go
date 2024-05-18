package main

import (
	"encoding/json"
	"log"

	"github.com/valyala/fastjson"
)

func main() {
	user := &User{ID: "test_value", Tags: []string{"1", "3", "2"}}
	// var user User
	data, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}

	log.Println("std json marshal", string(data))

	pool := &fastjson.ArenaPool{}
	arena := pool.Get()

	val := GetUserVal(arena, user)
	log.Println("manual build val", val.String())

	// Don't forget reset arena after String()
	arena.Reset()

	arena = pool.Get()

	new_val := GetReflectUserVal(arena, user)
	log.Println("reflect build val", new_val.String())

	// Don't forget reset arena after String()
	arena.Reset()
}
