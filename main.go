package main

import (
	"embed"
	"hatt/helpers"
	"hatt/variables"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var frontendAssets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	appOptions := &options.App{
		Title: "hatt",
		AssetServer: &assetserver.Options{
			Assets: frontendAssets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			&helpers.Helper{},
		},
	}

	if variables.MODE == "cli" {
		// hide app's window so it doesn't popup everytime it is recompiled
		appOptions.Width = 1
		appOptions.Height = 1
	} else {
		appOptions.WindowStartState = options.Maximised
	}

	// Create application with options
	err := wails.Run(appOptions)

	if err != nil {
		println("Error:", err.Error())
	}

}
