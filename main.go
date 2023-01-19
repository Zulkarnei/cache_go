package main

import (
	"fmt"
	"golang-ninja/basic/cache"
	"time"
)

func main() {
	c := cache.New()
	c.Set("key1", "value1", time.Second*5)
	c.Set("key2", "value2", time.Second*10)

	val, ok := c.Get("key1")
	if ok {
		fmt.Println("key1:", val)
	} else {
		fmt.Println("key1 not found")
	}

	val, ok = c.Get("key2")
	if ok {
		fmt.Println("key2:", val)
	} else {
		fmt.Println("key2 not found")
	}

	// Wait for key1 to expire
	time.Sleep(time.Second * 6)

	val, ok = c.Get("key1")
	if ok {
		fmt.Println("key1:", val)
	} else {
		fmt.Println("key1 not found")
	}
}
