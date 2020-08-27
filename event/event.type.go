package event

import (
	"github.com/bigobject-inc/golib/space"
)

// Event store the snapshot metadata
type Event struct {
	ID        string
	serial    int
	timestamp Timestamp
	frame     []Frame
}

// Timestamp store different representations of the event time
type Timestamp struct {
	integer int
	float   float64
	//datetime string // 存 2020-05-06T04:33:35.739Z 格式的, 要用自己轉
}

// Frame store a sensor and objects from it
type Frame struct {
	sensor   Sensor
	object   []Object
	gIDIndex map[string]int // $$ map[gID]index <-- ReID 可以順便建, 方便之後worker 可以直接用
	tIDIndex map[int]int    // $$ map[tID]index <--方便之後worker 可以直接用
}

// Sensor store serial and UUID of the sensor
type Sensor struct {
	ID     string
	serial int
}

// Object store all attributes and assignment to an object in a frame
type Object struct {
	name      string
	tID       int
	gID       string
	tag       Tag
	candidate []BackObject

	coordinate space.Point
	geohash    string

	boundingBox space.BoundingBox
	probability float64
}

// SimpleObject store an object with timestamp, sensor, geohash
type SimpleObject struct {
	ts      float64
	sensor  int
	name    string
	tID     int
	gID     string
	geohash string
}

// BackObject store an object with backtracking link count
type BackObject struct {
	sObj       SimpleObject
	confidence float64
	backLink   string
	count      int
}

// Tag store all tags assigned to an object
type Tag struct {
	pose         []string
	typedFeature map[string]*Feature
	wearable     []Wearable
	//POI	[]string
	//compound []string
	action []string
}

// Feature store generic data structure of a feature
type Feature struct {
	dataType string // $$ 這個欄位決定用什麼方式 parse 底下的 value
	value    interface{}
}

// Wearable define tangible relation of an object (to another object like person)
type Wearable struct {
	tID      int
	name     string
	relation string
}
