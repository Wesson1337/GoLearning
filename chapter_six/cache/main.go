package main

import (
	"fmt"
	"sync"
)

func main() {
	cache.mapping["key"] = "value"
	fmt.Println(Lookup("key"))
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
