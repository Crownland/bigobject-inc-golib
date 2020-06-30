package golib

import (
	"github.com/CloudChen2018/golib/irisweb"
	"github.com/CloudChen2018/golib/logger"
	"github.com/CloudChen2018/golib/miscellaneous"
)

// GoLibrary golang library service
type GoLibrary interface {
	GetServMiscellaneous() miscellaneous.Miscellaneous
	NewLogger(loggerFile, loggerLevel string) (logger.Logger, logger.Defer, error)
	NewRotationLogger(loggerFile, loggerLevel string) (logger.Logger, error)
	NewWeb(l logger.Logger, conf irisweb.Configure) (irisweb.Web, error)
}
