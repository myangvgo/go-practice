package main

import (
	"context"
	"net/http"
	"os"
	"syscall"
)

type Option func(o *options)

type options struct {
	ctx  context.Context
	sigs []os.Signal
}

func Signal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// lifecycle manager
type App struct {
	opts   options
	ctx    context.Context
	cancel func()
	server *http.Server
}

func New(opts ...Option) *App {
	options := options{
		ctx:  context.Background(),
		sigs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		opts:   options,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop() error {
	return nil
}

func main() {

}
