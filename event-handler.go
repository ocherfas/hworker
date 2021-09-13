package main

import (
	"log"

	hook "github.com/robotn/gohook"
)

type EventHandler struct {
	ActivityMonitors []*ActivityMonitor
	Enabled          bool
}

func NewEventHandler(activityMonitors []*ActivityMonitor) *EventHandler {
	return &EventHandler{
		ActivityMonitors: activityMonitors,
		Enabled:          false,
	}
}

func (eventHandler *EventHandler) toggle() {
	if eventHandler.Enabled {
		log.Println("End activity monitor")

		hook.End()
		eventHandler.Enabled = false
		return
	} else {
		log.Println("Start activity monitoring")

		for _, v := range eventHandler.ActivityMonitors {
			v.StartMonitor()
		}

		hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
			eventHandler.newEvent()
		})
		hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
			eventHandler.newEvent()
		})

		s := hook.Start()
		hook.Process(s)

		eventHandler.Enabled = true
	}
}

func (eventHandler *EventHandler) end() {
	if eventHandler.Enabled {
		hook.End()
	}
}

func (eventHandler *EventHandler) newEvent() {
	for _, v := range eventHandler.ActivityMonitors {
		err := v.newEvent()
		if err != nil {
			panic(err)
		}
	}
}
