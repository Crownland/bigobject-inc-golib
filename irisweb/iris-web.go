package irisweb

import (
	"github.com/kataras/iris/v12"
)

// Configure Web configure
type Configure struct {
	Host string             `json:"host"`
	Port int                `json:"port"`
	Iris iris.Configuration `json:"iris"`
}

// Paths Web paths
type Paths struct {
	Root        string
	WebRoot     string
	StaticPaths map[string]StaticPath
}

// Route Web route
type Route struct {
	Name   string
	Path   string
	Method string
	Handle []iris.Handler
}

// StaticFile Web static file
type StaticFile struct {
	Pathname string `json:"path"`
	Filename string `json:"file"`
	Absolute string `json:"absolute"`
	Relative string `json:"relative"`
}

// StaticPath Web AR(absolute and relative)Path
type StaticPath struct {
	Absolute string `json:"absolute"`
	Relative string `json:"relative"`
}

// Web Web service
type Web interface {
	CatchAPIError(ctx iris.Context, ctrlName string)
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
	ThrowAPIError(statucCode int, message string)
	ThrowAPIErrorf(statucCode int, format string, a ...interface{})
	Start() error
	Stop() error
}