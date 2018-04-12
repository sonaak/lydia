package app

import (
	"github.com/evilwire/go-env"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"sync"
	"time"
)

type MetaData struct {
	Ghash     string    `env:"GHASH" json:"ghash"`
	BuildTime time.Time `env:"BUILD_TIME" json:"build-time"`
	VERSION   string    `env:"VERSION" json:"version"`
}

type App struct {
	MetaData  *MetaData
	web       *echo.Echo
	StartTime time.Time
}

func NewApp(reader goenv.EnvReader) (*App, error) {
	app := App{
		web:      echo.New(),
		MetaData: &MetaData{},
	}
	marshaler := goenv.DefaultEnvMarshaler{reader}

	err := marshaler.Unmarshal(app.MetaData)
	if err != nil {
		return nil, err
	}

	app.web.Use(middleware.Logger())
	app.web.GET("/healthcheck", app.HealthCheck)

	return &app, nil
}

func (app *App) HealthCheck(context echo.Context) error {
	return context.JSON(200, app.MetaData)
}

func (app *App) Run() error {
	wg := sync.WaitGroup{}
	errStream := make(chan error)
	app.StartTime = time.Now()

	wg.Add(1)
	go func() {
		defer wg.Done()

		errStream <- app.web.Start(":9000")
	}()

	wg.Wait()
	return <-errStream
}
