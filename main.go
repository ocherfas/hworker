package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	hook "github.com/robotn/gohook"
)

const secondsActivicy = 15
const activityMax = secondsActivicy * time.Second
const inactivityTime = 5 * time.Second

func main() {
	fmt.Println("Running...")

	activityStart := time.Now()
	lastActionTime := time.Now()

	newEvent := func() {
		currentTime := time.Now()
		if currentTime.Sub(lastActionTime) >= inactivityTime {

			activityStart = currentTime
			lastActionTime = currentTime
			return
		} else if currentTime.Sub(activityStart) >= activityMax {

			message := fmt.Sprintf("You worked for %d already! You should take a break.\nLet your eyes rest for a bit :)", secondsActivicy)
			err := beeep.Notify("Take a break!", message, "")
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
