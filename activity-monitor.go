package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

type ActivityMonitor struct {
	inactivityTime  time.Duration
	maxActivityTime time.Duration
	messageFormat   string
	activityStart   time.Time
	lastActionTime  time.Time
}

func NewActivityMonitor(inactivityTime time.Duration, maxActivityTime time.Duration, messageFormat string) ActivityMonitor {
	return ActivityMonitor{
		messageFormat:   messageFormat,
		inactivityTime:  inactivityTime,
		maxActivityTime: maxActivityTime,
	}
}

func (activityMonitor *ActivityMonitor) StartMonitor() {
	current := time.Now()
	activityMonitor.activityStart = current
	activityMonitor.lastActionTime = current
}

func (activityMonitor *ActivityMonitor) newEvent() error {
	currentTime := time.Now()

	if currentTime.Sub(activityMonitor.lastActionTime) >= activityMonitor.inactivityTime {
		activityMonitor.activityStart = currentTime
		activityMonitor.lastActionTime = currentTime
		return nil
	} else if currentTime.Sub(activityMonitor.activityStart) >= activityMonitor.maxActivityTime {
		message := fmt.Sprintf(activityMonitor.messageFormat, activityMonitor.maxActivityTime)
		err := beeep.Alert("Take a break!", message, "")

		activityMonitor.activityStart = currentTime
		activityMonitor.lastActionTime = currentTime
		return err
	} else {
		activityMonitor.lastActionTime = currentTime
		return nil
	}
}
