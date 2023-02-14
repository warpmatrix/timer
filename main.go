package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	name := os.Args[1]
	fmt.Println("new task:", name)
	count_down(45*time.Minute, func() {
		Notify("session end", "it's time to take a break")
	})
	count_down(15*time.Minute, func() {
		Notify("time's up", "the rest time is finished")
	})
}

func count_down(duration time.Duration, handler func()) {
	seconds := int(duration.Seconds())
	tick := time.Tick(time.Second)
	for count := seconds; count >= 0; count-- {
		fmt.Printf("\rcount down: %02v:%02v", count/60, count%60)
		<-tick
	}
	fmt.Println()
	handler()
}
