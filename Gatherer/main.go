package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(get_data)
	s.StartAsync()
	s.StartBlocking()
}

func get_data() {
	fmt.Println("getting data")
}
