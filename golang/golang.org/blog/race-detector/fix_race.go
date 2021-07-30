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
	reset := make(chan bool)
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		reset <- true
	})
	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}
