package main

import (
	"encoding/json"
	"log"

	"github.com/valyala/fastjson"
)

func main() {
	pool := &fastjson.ArenaPool{}
	// user := &User{}
	var user *User
	data, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}

	log.Println("std marshal", string(data))

	arena := pool.Get()

	val := GetUserVal(arena, user)
	log.Println("manual val", val.String())

	// Don't forget reset arena after String()
	arena.Reset()

	gUser := &GenerateUser{}
	new_val := MakeGenerateUserVal(arena, gUser)
	log.Println("generated val", new_val.String())

	// Don't forget reset arena after String()
	arena.Reset()
}
