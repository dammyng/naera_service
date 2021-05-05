package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
)

func main()  {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(guess)
	s.StartAsync()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func guess()  {
	fmt.Println(rand.Intn(100))
}