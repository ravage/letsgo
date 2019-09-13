package handlers

import "letsgo/cmd/web/config"

// HandlerContext store shared application context across handlers
type HandlerContext struct {
	app *config.Application
}

// NewHandlerContext create instance
func NewHandlerContext(app *config.Application) *HandlerContext {
	handlerContext := HandlerContext{app: app}
	return &handlerContext
}
