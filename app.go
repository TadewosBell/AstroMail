package main

import (
	emailparser "AstroMail/email-parser"
	"AstroMail/storage"
	"context"
	"fmt"
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

// Launch SMTP Server
func (a *App) Launch_Smtp_Server(domain string, aws_id string, aws_secret string) {
	fmt.Println(domain, aws_id, aws_secret)
	storage.SaveCredentials(domain, aws_id, aws_secret)
}

// Greet returns a greeting for the given name
func (a *App) Get_Inbox() string {
	email, err := emailparser.ParseEmail("sample.eml")
	if err != nil {
		return "Error"
	}

	return email
}
