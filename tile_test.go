package tbox

import (
	"testing"
)

func isIn(x int, l []int) bool {
	for _, b := range l {
		if b == x {
			return true
		}
	}
	return false
}

func TestTileChildren(t *testing.T) {
	tile := Tile{Z: 0, X: 0, Y: 0}
	expectedX := []int{tile.X * 2, tile.X*2 + 1}
	expectedY := []int{tile.Y * 2, tile.Y*2 + 1}

	children := tile.Children()

	if len(children) != 4 {
		t.Fatalf("each tile has 4 children")
	}

	for _, e := range children {
		if e.Z != tile.Z+1 {
			t.Fatalf("child zoom must be tile zoom + 1")
		}

		if !isIn(e.X, expectedX) {
			t.Fatalf("child X must be tile X*2 or X*2+1")
		}

		if !isIn(e.Y, expectedY) {
			t.Fatalf("child Y must be tile Y*2 or Y*2+1")
		}
	}
}

func TestTileParent(t *testing.T) {
	var tests = []struct {
		child    Tile
		expected Tile
	}{
		{child: Tile{Z: 1, X: 1, Y: 1}, expected: Tile{Z: 0, X: 0, Y: 0}},
		{child: Tile{Z: 1, X: 1, Y: 0}, expected: Tile{Z: 0, X: 0, Y: 0}},
		{child: Tile{Z: 1, X: 0, Y: 0}, expected: Tile{Z: 0, X: 0, Y: 0}},
		{child: Tile{Z: 1, X: 0, Y: 1}, expected: Tile{Z: 0, X: 0, Y: 0}},
	}

	for _, test := range tests {
		got := test.child.Parent()
		if got != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, got)
		}
	}
}

func TestTileToBox(t *testing.T) {
	var tests = []struct {
		z      int
		x      int
		y      int
		minLng float64
		minLat float64
		maxLng float64
		maxLat float64
	}{
		{z: 11, x: 525, y: 761, minLng: -87.71484375, minLat: 41.77131167976407, maxLng: -87.5390625, maxLat: 41.9022770409637},
		{z: 15, x: 17599, y: 10756, minLng: 13.348388671875, minLat: 52.44931414086969, maxLng: 13.359375, maxLat: 52.456009392640745},
		{z: 11, x: 1095, y: 641, minLng: 12.48046875, minLat: 55.57834467218205, maxLng: 12.65625, maxLat: 55.67758441108952},
	}

	for _, test := range tests {
		tile := Tile{Z: test.z, X: test.x, Y: test.y}
		tbox := tile.ToBox()

		if tbox.MaxLat != test.maxLat || tbox.MaxLng != test.maxLng || tbox.MinLat != test.minLat || tbox.MinLng != test.minLng {
			t.Errorf("Expected: %v, got: %v", BoundingBox{MinLng: test.minLng, MinLat: test.minLat, MaxLng: test.maxLng, MaxLat: test.maxLat}, tbox)
		}
	}
}

func TestTileContainsPoint(t *testing.T) {
	var tests = []struct {
		tile     Tile
		lat      float64
		lng      float64
		expected bool
	}{
		{tile: Tile{Z: 11, X: 525, Y: 761}, lng: -87.65, lat: 41.84, expected: true},
		{tile: Tile{Z: 11, X: 1099, Y: 641}, lng: 12.568337, lat: 55.67609, expected: false},
	}

	for _, test := range tests {
		p := Point{Lng: test.lng, Lat: test.lat}
		i, _ := test.tile.ContainsPoint(p)

		if i != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, i)
		}
	}

	tile2 := Tile{Z: 11, X: 1099, Y: 641}
	_, err := tile2.ContainsPoint(Point{Lng: 999, Lat: 999})
	expected := "Point{Lat: 999, Lng: 999} - invalid point"

	if err.Error() != expected {
		t.Errorf("Expected: %v, got: %v", expected, err.Error())
	}

}

func TestTileToPoint(t *testing.T) {
	var tests = []struct {
		tile Tile
		p    Point
	}{
		{tile: Tile{Z: 11, X: 525, Y: 761}, p: Point{Lng: -87.626953125, Lat: 41.83679436036388}},
		{tile: Tile{Z: 15, X: 17599, Y: 10756}, p: Point{13.3538818359375, 52.45266176675521}},
	}

	for _, test := range tests {
		p := test.tile.ToPoint()

		if p.Lat != test.p.Lat || p.Lng != test.p.Lng {
			t.Errorf("expected: %v, got: %v", test.p, p)
		}
	}
}
