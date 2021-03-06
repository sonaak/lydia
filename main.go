package main

import (
	"github.com/evilwire/go-env"
	"github.com/sonaak/lydia/app"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
)



func main() {
	env := goenv.NewOsEnvReader()

	application, err := app.NewApp(env)
	if err != nil {
		panic(err)
	}



	panic(application.Run())
}
