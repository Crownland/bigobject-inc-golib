package irisweb

import (
	"github.com/kataras/iris/v12"
)

// Web Web service
type Web interface {
	CatchHandleError(ctx iris.Context, handleName string)
	GetIris() *iris.Application
	GetPaths() Paths
	GetStaticFileString(pathname, filename string) string
	GetStaticPaths(name string) (StaticPath, error)
	RemoveFileFromStaticFileString(fileString string) error
	SetRoutes(routes []Route) error
	SetStaticPaths(name string, path StaticPath, options iris.DirOptions) error
	StaticFileStringToFile(fileString string) (StaticFile, error)
	StaticFileToString(file StaticFile) string
	StaticFileToURL(file StaticFile) string
	ThrowHandleError(statucCode int, params HandleErrorParams)
	Start() error
	Stop() error
}
