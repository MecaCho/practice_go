package gotime

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func CronFunc() {
	c := cron.New()
	c.AddFunc("* * * * * *", RunEverySecond)
	// go func() {
	// }()
	c.Start()
	defer c.Stop()

	// sig := make(chan os.Signal)
	// signal.Notify(sig, os.Interrupt, os.Kill)
	// sig := make(chan int)
	time.Sleep(5 * time.Second)
	// <-sig
}

func RunEverySecond() {
	fmt.Printf("%v\n", time.Now())
}
