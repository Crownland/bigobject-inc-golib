package golib

import (
	"github.com/bigobject-inc/golib/irisweb"
	"github.com/bigobject-inc/golib/logger"
	"github.com/bigobject-inc/golib/miscellaneous"
)

// GoLibrary golang library service
type GoLibrary interface {
	GetServMiscellaneous() miscellaneous.Miscellaneous
	GetVersion() string
	NewLogger(loggerFile, loggerLevel string) (logger.Logger, logger.Defer, error)
	NewRotationLogger(loggerFile, loggerLevel string) (logger.Logger, error)
	NewWeb(l logger.Logger, conf irisweb.Configure, rootPath string) (irisweb.Web, error)
	NewWebController(path, method, name, summary string) (irisweb.Controller, error)
}
