package main

import (
	"time"

	"github.com/roman-baldaev/mit-distributed-systems/mr"
)

func main() {
	mr.MakeMaster([]string{"ss"}, 4)

	// Your code here.
	for i := 0; i < 1000; i++ {
		go mr.CallExample()
	}
	time.Sleep(time.Second * 3)
}
