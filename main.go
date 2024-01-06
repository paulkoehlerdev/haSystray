package main

import (
	"bytes"
	"fmt"
	"github.com/getlantern/systray"
	"haSystray/icons"
	"log/slog"
	"net/http"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	config := ReadConfigFromFile()

	systray.SetIcon(icons.GetDefaultIcon())
	systray.SetTooltip(config.ProgramTooltip)

	addActionsToMenu(config.Actions, nil)
	handleQuitButton()
}

func addActionsToMenu(actions []Action, groupMenuItem *systray.MenuItem) {
	for _, action := range actions {
		addActionToMenu(action, groupMenuItem)
	}
}

func addActionToMenu(action Action, groupMenuItem *systray.MenuItem) {
	var menuItem *systray.MenuItem

	if groupMenuItem == nil {
		menuItem = systray.AddMenuItem(action.Name, action.Tooltip)
	} else {
		menuItem = groupMenuItem.AddSubMenuItem(action.Name, action.Tooltip)
	}

	if len(action.Actions) > 0 {
		addActionsToMenu(action.Actions, menuItem)
	}

	if action.Type != "" && action.Webhook != "" {
		go func() {
			for range menuItem.ClickedCh {
				handleActionButton(menuItem, action)
			}
		}()
	} else if len(action.Actions) == 0 {
		menuItem.Disable()
	}
}

func handleActionButton(menuItem *systray.MenuItem, action Action) {
	menuItem.Disable()
	defer menuItem.Enable()

	body := bytes.NewReader(action.Data)

	slog.Info(fmt.Sprintf("making request: %s", action.Webhook), "action", action.Name, "url", action.Webhook, "type", action.Type, "data", string(action.Data))

	request, err := http.NewRequest(action.Type, action.Webhook, body)
	if err != nil {
		slog.Error(fmt.Sprintf("error while making request: %s", err.Error()), "error", err.Error(), "action", action.Name, "url", action.Webhook)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(request)
	if err != nil {
		slog.Error(fmt.Sprintf("error in request: %s", err.Error()), "error", err.Error(), "action", action.Name, "url", action.Webhook)
	}
}

func handleQuitButton() {
	mQuit := systray.AddMenuItem("Quit", "Quit")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	slog.Info("exiting")
}
