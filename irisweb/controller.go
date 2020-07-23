package irisweb

import (
	"github.com/kataras/iris/v12"
)

// Controller controller
type Controller interface {
	AddHandle(handle iris.Handler)
	ClearHandle()
	GetHandles() []iris.Handler
	GetInfo() ControllerInfo
}
