package cctv

// Detect CCTV detect view range
type Detect struct {
	Area       [][]float64 //closed polygon
	Distance   float64     //detecting distance
	Azimuth    float64     //detecting direction
	StartAngle float64     //sector's start azimuth
	EndAngle   float64     //sector's end azimuth
}

// Point World coordinates
type Point struct {
	Longitude float64 //x
	Latitude  float64 //y
	Height    float64 //z
}

// Pixel Image pixel coordinates
type Pixel struct {
	U float64 //x
	V float64 //y
}
