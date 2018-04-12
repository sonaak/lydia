package main

import (
	"github.com/pkg/errors"
	"time"
	"github.com/evilwire/go-env"
)

type App struct {}


func (app *App) Run() error {
	for {
		time.Sleep(1 * time.Second)
	}
	return errors.New("Broken")
}


func MustSetup(env goenv.EnvReader) *App {
	return &App{}
}

func main() {
	env := goenv.NewOsEnvReader()
	app := MustSetup(env)

	panic(app.Run())
}
