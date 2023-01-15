package main

import (
	"os"
	"time"

	"github.com/adrianwudev/learngowithtests/countdown"
)

func main() {
	// sleeper := &countdown.ConfigurableSleeper{1 * time.Second, time.Sleep}
	// in order to create an instance of ConfigurableSleeper with private properties,
	// we create a function with parameters to create the instance.
	sleeper := countdown.NewConfigurableSleeper(1*time.Second, time.Sleep)
	countdown.CountDown(os.Stdout, sleeper)
}
