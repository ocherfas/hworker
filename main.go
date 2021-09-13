package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
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

const ENABLE_TEXT = "Enable"
const DISABLE_TEXT = "Disable"

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

	eventHandler := NewEventHandler(activityMonitors)

	systray.Run(func() {
		systray.SetTooltip("Healthy Worker")
		systray.SetTitle("Healthy Worker")

		mEnable := systray.AddMenuItem("", "")
		mQuit := systray.AddMenuItem("Quit", "Quit")

		toggleAndSet := func() {
			eventHandler.toggle()
			if eventHandler.Enabled {
				mEnable.SetTitle(DISABLE_TEXT)
				mEnable.SetTooltip(DISABLE_TEXT)
			} else {
				mEnable.SetTitle(ENABLE_TEXT)
				mEnable.SetTooltip(ENABLE_TEXT)
			}
		}

		toggleAndSet()

		go func() {
			for {
				select {
				case <-mQuit.ClickedCh:
					systray.Quit()
					return
				case <-mEnable.ClickedCh:
					toggleAndSet()
				}
			}

		}()
	}, func() {
		eventHandler.end()
	})
}
