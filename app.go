package main

import (
	"context"
	"flag"
	"fmt"
	"hatt/variables"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {

	flag.Parse()
	fmt.Println("args : ", flag.Args())
	for index, arg := range flag.Args() {
		if index == 0 {
			variables.MODE = arg
		}
	}

	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
