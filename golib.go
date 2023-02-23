package golib

import (
	"github.com/bigobject-inc/golib/logger"
	"github.com/bigobject-inc/golib/miscellaneous"
	"github.com/bigobject-inc/golib/space"
	"github.com/bigobject-inc/golib/space/cctv"
)

// GoLibrary golang library service
type GoLibrary interface {
	GetServMiscellaneous() miscellaneous.Miscellaneous
	GetVersion() string
	NewLogger(loggerFile, loggerLevel string) (logger.Logger, logger.Defer, error)
	NewRotationLogger(loggerFile, loggerLevel string) (logger.Logger, error)
	GetServGeometry() space.Geometry
	NewCCTV(id string, geolocation [2]float64, resolution []interface{}, cameraMatrix [][]float64, rmat [][]float64, tvec [][]float64, dist [][]float64) cctv.CCTV
}
