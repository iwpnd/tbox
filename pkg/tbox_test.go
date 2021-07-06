package tbox

import (
	"testing"
)

func TestInvalidPointError(t *testing.T) {
	expected := "Point{Lat: 52, Lng: -190} - invalid point"
	p := Point{Lat: 52, Lng: -190}

	_, err := p.ToTile(15)

	if err.Error() != expected {
		t.Errorf("Expected: %v, got: %v", expected, err)

	}

}

func TestPointValid(t *testing.T) {
	tests := []struct {
		lat, lng float64
		expected bool
	}{
		{lat: 52, lng: 13, expected: true},
		{lat: 91, lng: 13, expected: false},
		{lat: 52, lng: 181, expected: false},
		{lat: -91, lng: -181, expected: false},
	}

	for _, test := range tests {
		got := Point{Lat: test.lat, Lng: test.lng}.Valid()

		if got != test.expected {
			t.Errorf("Expected %v, got: %v", test.expected, got)
		}
	}
}

func TestPointToTile(t *testing.T) {
	var tests = []struct {
		lat float64
		lng float64
		z   int
		x   int
		y   int
	}{
		{lat: 41.84, lng: -87.65, z: 3, x: 2, y: 2},
		{lat: 52.44950563632098, lng: 13.357951727129988, z: 15, x: 17599, y: 10756},
	}

	var p Point
	var tile Tile

	for _, test := range tests {
		p = Point{Lat: test.lat, Lng: test.lng}
		tile, _ = p.ToTile(test.z)

		if tile.Z != test.z || tile.X != test.x || tile.Y != test.y {
			t.Errorf("Expected: %v, got: %v", Tile{test.z, test.x, test.y}, tile)
		}
	}
}

func TestPointInTile(t *testing.T) {
	var tests = []struct {
		lat      float64
		lng      float64
		z        int
		x        int
		y        int
		expected bool
	}{
		{lat: 41.84, lng: -87.65, z: 3, x: 2, y: 2, expected: true},
		{lat: 52.44950563632098, lng: 13.357951727129988, z: 15, x: 17599, y: 10756, expected: true},
		{lat: 55.676098, lng: 12.568337, z: 11, x: 1095, y: 641, expected: true},
		{lat: 52.25, lng: 13.37, z: 11, x: 111, y: 111, expected: false},
	}

	var p Point
	var tile Tile
	var i bool

	for _, test := range tests {
		tile = Tile{Z: test.z, X: test.x, Y: test.y}
		p = Point{Lat: test.lat, Lng: test.lng}
		i, _ = p.InTile(tile)

		if i != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, i)
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

	var tile Tile
	var tbox Tilebox

	for _, test := range tests {
		tile = Tile{Z: test.z, X: test.x, Y: test.y}
		tbox = tile.ToBox()

		if tbox.MaxLat != test.maxLat || tbox.MaxLng != test.maxLng || tbox.MinLat != test.minLat || tbox.MinLng != test.minLng {
			t.Errorf("Expected: %v, got: %v", Tilebox{MinLng: test.minLng, MinLat: test.minLat, MaxLng: test.maxLng, MaxLat: test.maxLat}, tbox)
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
	}

	var p Point

	for _, test := range tests {
		p = Point{Lng: test.lng, Lat: test.lat}
		i, err := test.tile.ContainsPoint(p)

		if err != nil {
			t.Error("something went wrong")
		}

		if i != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, i)
		}
	}

}
