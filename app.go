package main

import (
	"context"
	"time"

	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

	go func() {
		var count uint = 0
		hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
			count++
		})
		hook.Process(hook.Start())
		for {
			time.Sleep(1 * time.Second)
			runtime.EventsEmit(ctx, "MouseRate", count)
			count = 0
		}
	}()
}
