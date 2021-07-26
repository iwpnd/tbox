package tbox

import (
	"fmt"
	"math"
)

type invalidPointError struct {
	p   Point
	msg string
}

// Tilebox ...
type Tilebox struct {
	MinLng, MinLat, MaxLng, MaxLat float64
}

// Tile ...
type Tile struct {
	Z, X, Y int
}

// Point ...
type Point struct {
	Lng, Lat float64
}

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

func (e *invalidPointError) Error() string {
	return fmt.Sprintf("Point{Lat: %v, Lng: %v} - %s", e.p.Lat, e.p.Lng, e.msg)
}

// Valid ...
func (p Point) Valid() bool {
	return p.Lng >= -180 && p.Lng <= 180 && p.Lat >= -90 && p.Lat <= 90
}

// ToTile ...
func (p Point) ToTile(z int) (Tile, error) {
	if !p.Valid() {
		return Tile{}, &invalidPointError{p: p, msg: "invalid point"}
	}

	latRad := degreeToRad(p.Lat)
	n := math.Pow(2, float64(z))

	xtile := int((p.Lng + 180) / 360 * n)
	ytile := int((1 - math.Asinh(math.Tan(latRad))/math.Pi) / 2 * n)

	return Tile{z, xtile, ytile}, nil
}

// InTile ...
func (p Point) InTile(t Tile) (bool, error) {
	if !p.Valid() {
		return false, &invalidPointError{p: p, msg: "invalid point"}
	}
	tbox := t.ToBox()

	return p.Lng > tbox.MinLng && p.Lat > tbox.MinLat && p.Lng < tbox.MaxLng && p.Lat < tbox.MaxLat, nil
}

// ToBox ...
func (t Tile) ToBox() Tilebox {
	return Tilebox{
		MaxLat: tileToLat(t.Y, t.Z),
		MinLng: tileToLng(t.X, t.Z),
		MinLat: tileToLat(t.Y+1, t.Z),
		MaxLng: tileToLng(t.X+1, t.Z),
	}
}

// ContainsPoint ...
func (t Tile) ContainsPoint(p Point) (bool, error) {
	if !p.Valid() {
		return false, &invalidPointError{p: p, msg: "invalid point"}
	}
	tbox := t.ToBox()

	return p.Lat > tbox.MinLat && p.Lat < tbox.MaxLat && p.Lng > tbox.MinLng && p.Lng < tbox.MaxLng, nil
}

// ToPoint ...
func (t Tile) ToPoint() Point {
	tbox := t.ToBox()

	cLng := tbox.MinLng + (tbox.MaxLng-tbox.MinLng)/2
	cLat := tbox.MinLat + (tbox.MaxLat-tbox.MinLat)/2

	return Point{Lng: cLng, Lat: cLat}
}
