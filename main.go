package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	hook "github.com/robotn/gohook"
)

const maxActivityTimeString = "15s"
const inactivityTimeString = "5s"

func main() {
	fmt.Println("Running...")

	maxTimeActivity, errParse1 := time.ParseDuration(maxActivityTimeString)
	if errParse1 != nil {
		err := fmt.Errorf("Incorrect format for max time activity")
		panic(err)
	}
	inactivityTime, errParse2 := time.ParseDuration(inactivityTimeString)
	if errParse2 != nil {
		err := fmt.Errorf("Incorrect format for inactivity time")
		panic(err)
	}

	activityStart := time.Now()
	lastActionTime := time.Now()

	newEvent := func() {
		currentTime := time.Now()
		if currentTime.Sub(lastActionTime) >= inactivityTime {
			activityStart = currentTime
			lastActionTime = currentTime
			return
		} else if currentTime.Sub(activityStart) >= maxTimeActivity {
			message := fmt.Sprintf("You worked for %s already! You should take a break.\nLet your eyes rest for a bit :)", maxTimeActivity)
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
