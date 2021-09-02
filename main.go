package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	hook "github.com/robotn/gohook"
)

func main() {
	config := NewConfig()
	err := config.readConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("Running...")

	activityStart := time.Now()
	lastActionTime := time.Now()

	newEvent := func() {
		currentTime := time.Now()
		if currentTime.Sub(lastActionTime) >= config.inactivityTime {
			activityStart = currentTime
			lastActionTime = currentTime
			return
		} else if currentTime.Sub(activityStart) >= config.maxActivityTime {
			message := fmt.Sprintf("You worked for %s already! You should take a break.\nLet your eyes rest for a bit :)", config.maxActivityTime)
			err := beeep.Alert("Take a break!", message, "")
			if err != nil {
				panic(err)
			}

			activityStart = currentTime
			lastActionTime = currentTime
			return
		} else {
			lastActionTime = currentTime
			return
		}
	}

	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		newEvent()
	})
	hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
		newEvent()
	})

	s := hook.Start()
	<-hook.Process(s)
}
