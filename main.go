package main

import (
	"fmt"
	"os"
	"time"

	"github.com/getlantern/systray"
)

var (
	timezone string
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	timezone = "Local"

	systray.SetIcon(getIcon("assets/clock.ico"))

	localTime := systray.AddMenuItem("Local time", "Local time")
	hcmcTime := systray.AddMenuItem("Ho Chi Minh time", "Asia/Ho_Chi_Minh")
	sydTime := systray.AddMenuItem("Sydney time", "Australia/Sydney")
	gdlTime := systray.AddMenuItem("Guadalajara time", "America/Mexico_City")
	sfTime := systray.AddMenuItem("San Fransisco time", "America/Los_Angeles")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			systray.SetTitle(getClockTime(timezone))
			systray.SetTooltip(timezone + " timezone")
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-localTime.ClickedCh:
				timezone = "Local"
			case <-hcmcTime.ClickedCh:
				timezone = "Asia/Ho_Chi_Minh"
			case <-sydTime.ClickedCh:
				timezone = "Australia/Sydney"
			case <-gdlTime.ClickedCh:
				timezone = "America/Mexico_City"
			case <-sfTime.ClickedCh:
				timezone = "America/Los_Angeles"
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleaning stuff here.
}

func getClockTime(tz string) string {
	t := time.Now()
	utc, _ := time.LoadLocation(tz)

	return t.In(utc).Format("15:04:05")
}

// getIcon reads an icon file from the given path.
func getIcon(filePath string) []byte {
	icon, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error during downloading icon: %v\n", err)
	}
	return icon
}
