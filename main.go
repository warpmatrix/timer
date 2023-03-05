package main

import (
	"fmt"
	"time"
	"timer/utils"

	flag "github.com/spf13/pflag"
)

func main() {
	name := "untitled"
	if len(flag.Args()) > 0 {
		name = flag.Args()[0]
	}
	fmt.Println("new task:", name)
	sess, _ := time.ParseDuration(utils.Get_flag("time"))
	rest, _ := time.ParseDuration(utils.Get_flag("rest"))
	count_down(sess, func() {
		utils.Notify("session end", "it's time to take a break")
	})
	count_down(rest, func() {
		utils.Notify("time's up", "the rest time is finished")
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
