package space

// Geometry space geometry service
type Geometry interface {
	Encode(longitude float64, latitude float64, precision int) string
	Decode(geohash string) BoundingBox
	Adjacent(geohash string, direction string) string
	GeohashNeighbors(geohash string) []string
	GeohashDirection(geohash1 string, geohash2 string) string
	PolygonBoundingBox(polygon [][]float64) (rectangle [][]float64)
	GeohashWithinArea(polygon [][]float64, precision int) (geohash []string)
	FindNearestPoints(x, y float64, pointlist [][]float64) Point
	GeoDistance(lng1 float64, lat1 float64, lng2 float64, lat2 float64) (distance float64)
	GeoBearing(lng1 float64, lat1 float64, lng2 float64, lat2 float64) (bearing float64)
	GeoPathIntersection(lng1 float64, lat1 float64, bearing1 float64, lng2 float64, lat2 float64, bearing2 float64) (Point, error)
	GeoProjection(lng1 float64, lat1 float64, direction float64, distance float64) Point
	PointFromLatLng(lat, lng float64) Point
	PointInPolygon(point Point, polygon [][]float64) bool
}
