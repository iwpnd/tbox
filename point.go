package tbox

import (
	"fmt"
	"math"
)

type invalidPointError struct {
	p   Point
	msg string
}

// Point ...
type Point struct {
	Lng, Lat float64
}

func (e *invalidPointError) Error() string {
	return fmt.Sprintf("Point{Lat: %v, Lng: %v} - %s", e.p.Lat, e.p.Lng, e.msg)
}

// Valid validates a Point
func (p Point) Valid() bool {
	return p.Lng >= -180 && p.Lng <= 180 && p.Lat >= -90 && p.Lat <= 90
}

// ToTile calculates Tile coordinates Z/X/Y from a given Point
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

// InTile validates if Point is in a given Tile bounding box
func (p Point) InTile(t Tile) (bool, error) {
	if !p.Valid() {
		return false, &invalidPointError{p: p, msg: "invalid point"}
	}
	tbox := t.ToBox()

	return p.Lng > tbox.MinLng && p.Lat > tbox.MinLat && p.Lng < tbox.MaxLng && p.Lat < tbox.MaxLat, nil
}
