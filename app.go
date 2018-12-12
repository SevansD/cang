package cang

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"
)

type App struct {
	needShutdown bool
	turnedOff    bool
	services     map[string]Service
	ready        bool
	logger       log.Logger
	ctx          context.Context
	finish       func()
}

func (app *App) Start(options ...Option) *App {
	for _, option := range options {
		e := option(app)
		if e != nil {
			fmt.Printf("Some option can't executed. See error: %s", e.Error())
			panic("Can't execute application option")
		}
	}
	if app.ctx == nil {
		ctx, finish := context.WithCancel(context.Background())
		app.ctx, app.finish = ctx, finish
	}
	app.ready = true
	return app
}

func (app *App) Work() {
	if !app.ready {
		panic("try to work not ready application")
	}
	for name, service := range app.services {
		go func() {
			app.logger.Print("Starting", name)
			err := service.Start(app)
			if err != nil {
				app.logger.Fatal(err)
			}
		}()
	}
	for {
		select {
		case <-app.ctx.Done():
			app.logger.Print("Go to stop all")
			for _, service := range app.services {
				service.Stop()
			}
			app.logger.Print("application was stopped")
			return
		default:
			runtime.Gosched()
			time.Sleep(time.Second)
		}
	}
}

func (app *App) Stop() {
	if app.finish != nil {
		app.finish()
	}
}
