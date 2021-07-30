package main

// Link: https://blog.golang.org/race-detector

import (
	"fmt"
	"math/rand"
	"time"
)

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}
