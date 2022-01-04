package tbox

import (
	"reflect"
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

func inDeep(x Tile, y []Tile) bool {
	for i := range y {
		if reflect.DeepEqual(x, y[i]) {
			return true
		}
	}
	return false
}

func TestFromBounds(t *testing.T) {
	tests := []struct {
		bb       BoundingBox
		z        int
		expected []Tile
	}{
		{
			bb: BoundingBox{MinLng: 10.045, MinLat: 51.2114, MaxLng: 13.825, MaxLat: 53.575},
			z:  7,
			expected: []Tile{
				{X: 67, Y: 41, Z: 7},
				{X: 68, Y: 41, Z: 7},
				{X: 67, Y: 42, Z: 7},
				{X: 68, Y: 42, Z: 7},
			},
		},
		{
			bb:       BoundingBox{MinLng: 9.7061, MinLat: 53.3942, MaxLng: 10.3019, MaxLat: 53.763},
			z:        7,
			expected: []Tile{{X: 67, Y: 41, Z: 7}},
		},
		{
			bb: BoundingBox{MinLng: 178.65, MinLat: 70.81, MaxLng: -177.58, MaxLat: 71.6},
			z:  7,
			expected: []Tile{
				{X: 127, Y: 26, Z: 7},
				{X: 0, Y: 26, Z: 7},
				{X: 127, Y: 27, Z: 7},
				{X: 0, Y: 27, Z: 7},
			},
		},
	}

	for _, test := range tests {
		got, err := FromBounds(test.bb, test.z)
		if err != nil {
			t.Fatal("cannot get tiles from bounds")
		}

		if len(got) != len(test.expected) {
			t.Fatal("length should be equal")
		}

		for i := range got {
			if !inDeep(got[i], test.expected) {
				t.Fatal(got[i], "not in expected: ", test.expected)
			}
		}
	}
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

func TestTileToBbox(t *testing.T) {
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
		tbox := tile.Bbox()

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
		i, _ := test.tile.Contains(p)

		if i != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, i)
		}
	}

	tile2 := Tile{Z: 11, X: 1099, Y: 641}
	_, err := tile2.Contains(Point{Lng: 999, Lat: 999})
	expected := "Point{Lat: 999, Lng: 999} - invalid point"

	if err.Error() != expected {
		t.Errorf("Expected: %v, got: %v", expected, err.Error())
	}

}

func TestTileCenter(t *testing.T) {
	var tests = []struct {
		tile Tile
		p    Point
	}{
		{tile: Tile{Z: 11, X: 525, Y: 761}, p: Point{Lng: -87.626953125, Lat: 41.83679436036388}},
		{tile: Tile{Z: 15, X: 17599, Y: 10756}, p: Point{13.3538818359375, 52.45266176675521}},
	}

	for _, test := range tests {
		p := test.tile.Center()

		if p.Lat != test.p.Lat || p.Lng != test.p.Lng {
			t.Errorf("expected: %v, got: %v", test.p, p)
		}
	}
}
