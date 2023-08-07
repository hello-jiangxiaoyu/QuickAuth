package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar", time.Second)
	c.Set("baz", 42, cache.NoExpiration)

	time.Sleep(4 * time.Second)
	fmt.Println(c.Get("foo"))
	fmt.Println(c.Get("baz"))
}
