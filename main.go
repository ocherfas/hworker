package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
	hook "github.com/robotn/gohook"
)

func convertConfigToActivityMonitors(config Config) ([]*ActivityMonitor, error) {
	activityMonitors := []*ActivityMonitor{}
	for i, v := range config.ActivityMonitorConfigs {
		maxTimeActivity, errParse1 := time.ParseDuration(v.MaxActivityTime)
		if errParse1 != nil {
			err := fmt.Errorf("ActivityMonitor config number %d incorrect: Incorrect format for max time activity", i+1)
			return nil, err
		}
		inactivityTime, errParse2 := time.ParseDuration(v.InactivityTime)
		if errParse2 != nil {
			err := fmt.Errorf("ActivityMonitor config number %d incorrect: Incorrect format for max time activity", i+1)
			return nil, err
		}

		newActivityMonitor := NewActivityMonitor(inactivityTime, maxTimeActivity, v.MessageFormat)
		activityMonitors = append(activityMonitors, &newActivityMonitor)
	}

	return activityMonitors, nil
}

func main() {
	config := NewConfig()
	err := config.readConfig()

	if err != nil {
		panic(err)
	}

	activityMonitors, err := convertConfigToActivityMonitors(config)
	if err != nil {
		panic(err)
	}

	systray.Run(func() {
		systray.SetTooltip("Healty Worker")
		systray.SetTitle("Healty Worker")
		mQuit := systray.AddMenuItem("Quit", "Quit")

		for _, v := range activityMonitors {
			v.StartMonitor()
		}

		newEvent := func() {
			for _, v := range activityMonitors {
				err := v.newEvent()
				if err != nil {
					panic(err)
				}
			}
		}

		hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
			newEvent()
		})
		hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
			newEvent()
		})

		s := hook.Start()
		hook.Process(s)

		go func() {
			<-mQuit.ClickedCh
			systray.Quit()
		}()
	}, func() {
		hook.End()
	})
}
