package space

//BoundingBox Geohash bounding box coordinates
type BoundingBox struct {
	NE     Point
	SW     Point
	Center Point
}

// Point World coordinates
type Point struct {
	Longitude float64 //x
	Latitude  float64 //y
	Height    float64 //z
}
