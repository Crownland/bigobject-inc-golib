package cctv

// CCTV cctv
type CCTV interface {
	DetectRange(maxDetectDistance, degreeScale float64) Detect
	ImagePixelToWorldPoint(u, v, height float64) Point
	WorldPointToImagePixel(lng, lat, height float64) Pixel
}
