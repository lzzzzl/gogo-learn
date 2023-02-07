package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func task() {
	fmt.Printf("Hello World, %s\n", time.Now().String())
}

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(task)

	// strings parse to duration
	s.Every("5m").Do(task)

	s.Every(5).Days().Do(task)

	s.Every(1).Month(1, 2, 3).Do(task)

	// set time
	s.Every(1).Day().At("10:30").Do(task)

	// set multiple times
	s.Every(1).Day().At("10:30;08:00").Do(task)

	s.Every(1).Day().At("10:30").At("08:00").Do(task)

	// Schedule each last day of the month
	s.Every(1).MonthLastDay().Do(task)

	// or each last day of every other month
	s.Every(2).MonthLastDay().Do(task)

	// cron expressions supported
	s.Cron("*/1 * * * *").Do(task) // every minute

	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	s.StartAsync()
	s.StartBlocking()
}
