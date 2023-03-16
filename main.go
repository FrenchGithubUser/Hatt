package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var frontendAssets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// var appWidth int = 1600
	// var appHeight int = 1000
	// if variables.MODE == "cli" {
	// 	appWidth = 1
	// 	appHeight = 1
	// }

	// Create application with options
	err := wails.Run(&options.App{
		Title: "hatt",
		// Width:  appWidth,
		// Height: appHeight,
		WindowStartState: options.Maximised,
		AssetServer: &assetserver.Options{
			Assets: frontendAssets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
