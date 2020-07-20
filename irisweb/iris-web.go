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

// HandleErrorParams web handle error params
type HandleErrorParams struct {
	Code           string
	DisplayMessage string
	HandleName     string
	Message        string
	StatusCode     int
	Params         map[string]interface{}
}

// HandleErrorResponse web handle error response
type HandleErrorResponse struct {
	Code    string
	Message string
	Params  map[string]interface{}
}

// Paths Web paths
type Paths struct {
	Root        string
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
