package tbox

import "math"

func radToDegree(rad float64) float64 {
	return rad * 180 / math.Pi
}

func degreeToRad(degree float64) float64 {
	return degree * math.Pi / 180
}

func tileToLng(x, z int) float64 {
	return float64(x)/math.Pow(2.0, float64(z))*360.0 - 180
}

func tileToLat(y, z int) float64 {
	n := math.Pi - (2.0*math.Pi*float64(y))/math.Pow(2.0, float64(z))
	return radToDegree(math.Atan(math.Sinh(n)))
}
