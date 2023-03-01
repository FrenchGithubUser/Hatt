package main

import (
	"context"
	"hatt/variables"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {

	variables.InitVariables()

	// creates user config dir, ignores the error if the directory already exists
	_ = os.Mkdir(variables.USER_CONFIG_DIR, os.ModePerm)

	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
