package tbox

import (
	"math"
)

func min(x, y float64) float64 {
	if x > y {
		return y
	}

	return x
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}

	return y
}

// Tile ...
type Tile struct {
	Z, X, Y int
}

// BoundingBox ...
type BoundingBox struct {
	MinLng, MinLat, MaxLng, MaxLat float64
}

// Bbox returns the bounding box of a given Tile
func (t Tile) Bbox() BoundingBox {
	return BoundingBox{
		MaxLat: tileToLat(t.Y, t.Z),
		MinLng: tileToLng(t.X, t.Z),
		MinLat: tileToLat(t.Y+1, t.Z),
		MaxLng: tileToLng(t.X+1, t.Z),
	}
}

// FromBounds returns all tiles of a certain zoom level that intersect an input
// bounding box
func FromBounds(b BoundingBox, z int) ([]Tile, error) {
	Z := int(math.Pow(2, float64(z)))

	var bbs []BoundingBox

	if b.MinLng > b.MaxLng {
		bbw := BoundingBox{MaxLat: b.MaxLat, MaxLng: b.MaxLng, MinLng: -180.0, MinLat: b.MinLat}
		bbe := BoundingBox{MaxLat: b.MaxLat, MaxLng: 180.0, MinLng: b.MinLng, MinLat: b.MinLat}
		bbs = []BoundingBox{bbw, bbe}
	} else {
		bbs = []BoundingBox{b}
	}

	var tiles []Tile

	for _, bb := range bbs {
		minlng := max(-180.0, bb.MinLng)
		minlat := max(-85.0, bb.MinLat)
		maxlng := min(180.0, bb.MaxLng)
		maxlat := min(85.0, bb.MaxLat)

		ult, err := Point{Lng: minlng, Lat: maxlat}.ToTile(z)

		if err != nil {
			return nil, err
		}

		lrt, err := Point{Lng: maxlng, Lat: minlat}.ToTile(z)

		if err != nil {
			return nil, err
		}

		for i := ult.X; i <= lrt.X; i++ {
			for j := ult.Y; j <= lrt.Y; j++ {
				// ignore coordinates >= 2 ** zoom
				if i >= Z {
					continue
				}

				if j >= Z {
					continue
				}

				tile := Tile{X: i, Y: j, Z: z}
				tiles = append(tiles, tile)
			}
		}
	}

	return tiles, nil
}

// Contains valides whether a given Point is within a given Tile
func (t Tile) Contains(p Point) (bool, error) {
	if !p.Valid() {
		return false, &invalidPointError{p: p, msg: "invalid point"}
	}
	tbox := t.Bbox()

	return p.Lat > tbox.MinLat && p.Lat < tbox.MaxLat && p.Lng > tbox.MinLng && p.Lng < tbox.MaxLng, nil
}

// Center returns the centerpoint of a given Tile
func (t Tile) Center() Point {
	tbox := t.Bbox()

	cLng := tbox.MinLng + (tbox.MaxLng-tbox.MinLng)/2
	cLat := tbox.MinLat + (tbox.MaxLat-tbox.MinLat)/2

	return Point{Lng: cLng, Lat: cLat}
}

// Children returns the four children tiles of the input tile
func (t Tile) Children() [4]Tile {
	x := t.X
	y := t.Y
	z := t.Z

	return [4]Tile{
		{X: x * 2, Y: y * 2, Z: z + 1},
		{X: x*2 + 1, Y: y * 2, Z: z + 1},
		{X: x * 2, Y: y*2 + 1, Z: z + 1},
		{X: x*2 + 1, Y: y*2 + 1, Z: z + 1},
	}
}

// Parent returns the parent tile of the input tile
func (t Tile) Parent() Tile {
	x := t.X
	y := t.Y
	z := t.Z

	if z == 0 {
		return t
	}

	if x%2 == 0 && y%2 == 0 {
		return Tile{Z: z - 1, X: x / 2, Y: y / 2}
	} else if x%2 == 0 {
		return Tile{Z: z - 1, X: x / 2, Y: (y - 1) / 2}
	} else if x%2 != 0 && y%2 == 0 {
		return Tile{Z: z - 1, X: (x - 1) / 2, Y: y / 2}
	} else {
		return Tile{Z: z - 1, X: (x - 1) / 2, Y: (y - 1) / 2}
	}
}
