package main

import (
	"github.com/evilwire/go-env"
	"github.com/sonaak/lydia/app"
)



func main() {
	env := goenv.NewOsEnvReader()

	application, err := app.NewApp(env)
	if err != nil {
		panic(err)
	}

	panic(application.Run())
}
