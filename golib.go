package golib

import (
	"github.com/CloudChen2018/golib-definitions/irisweb"
	"github.com/CloudChen2018/golib-definitions/logger"
	"github.com/CloudChen2018/golib-definitions/miscellaneous"
)

// GoLibrary golang library service
type GoLibrary interface {
	GetServMiscellaneous() miscellaneous.Miscellaneous
	GetVersion() string
	NewLogger(loggerFile, loggerLevel string) (logger.Logger, logger.Defer, error)
	NewRotationLogger(loggerFile, loggerLevel string) (logger.Logger, error)
	NewWeb(l logger.Logger, conf irisweb.Configure, rootPath string) (irisweb.Web, error)
}
