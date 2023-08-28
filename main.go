package main

import (
	"fmt" // Delete this line to eliminate import from internal.
	"keyrehack/configs"
	"keyrehack/internal/keyboard" // Delete this line to eliminate import from internal.

	// Delete this line to eliminate import from internal.
	"time"

	"keyrehack/icon"

	"fyne.io/systray"
)

func main() {
	onExit := func() {
		now := time.Now()
		fmt.Println("Exit at", now.String())
	}

	systray.Run(onReady, onExit)
}

func addQuitItem() {
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
	systray.AddSeparator()
}

func onReady() {
	icon := icon.GetIcon()
	systray.SetTemplateIcon(icon, icon)
	systray.SetTitle("Keyrehack")
	systray.SetTooltip("Keyrehack")
	addQuitItem()
	subMenu := systray.AddMenuItem("Keymap Mode", "Keymap Mode")
	setting1 := subMenu.AddSubMenuItemCheckbox("setting 1", "setting 1", true)
	setting2 := subMenu.AddSubMenuItemCheckbox("setting 2", "setting 2", false)
	reset := systray.AddMenuItem("Reset", "Reset")

	simpleReplace1, combinationReplace1 := configs.GetKeySetting1()
	simpleReplace2, combinationReplace2 := configs.GetKeySetting2()

	go keyboard.SetHack(simpleReplace1, combinationReplace1)

	unckeckAllSetings := func() {
		setting1.Uncheck()
		setting2.Uncheck()
	}

	go func() {
		for {
			select {
			case <-setting1.ClickedCh:
			case <-reset.ClickedCh:
				unckeckAllSetings()
				setting1.Check()
				keyboard.UnsetHack()
				go keyboard.SetHack(simpleReplace1, combinationReplace1)
			case <-setting2.ClickedCh:
				unckeckAllSetings()
				setting2.Check()
				keyboard.UnsetHack()
				go keyboard.SetHack(simpleReplace2, combinationReplace2)
			}
		}
	}()
}
