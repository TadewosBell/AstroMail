package main

import (
	emailparser "changeme/email-parser"
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Get_Inbox() string {
	email, err := emailparser.ParseEmail("sample.eml")
	if err != nil {
		return "Errpr"
	}

	return email
}
