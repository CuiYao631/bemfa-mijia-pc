package app

import (
	"changeme/app/user"
	"context"
	"fmt"
	"log"
)

// App struct
type App struct {
	ctx context.Context
	u   user.User
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.u = user.NewUser()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RegisterByEmail(email, password string) {
	a.u.RegisterByEmail(a.ctx, email, password)
}
func (a *App) LoginByEmail(email, password string) error {
	log.Println("LoginByEmail")
	log.Println(email, password)
	err := a.u.LoginByEmail(a.ctx, email, password)
	if err != nil {
		return err
	}

	return nil
}
