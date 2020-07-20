package miscellaneous

// Miscellaneous Miscellaneous service
type Miscellaneous interface {
	CopyFile(src, dst string) (int64, error)
	CreateUUID() string
	FindInArray(val interface{}, array interface{}) (bool, int)
	GetAppRootPath() string
	InArray(val interface{}, array interface{}) bool
	StructToJSONByte(input interface{}) ([]byte, error)
	StructToJSONString(input interface{}) (string, error)
}
