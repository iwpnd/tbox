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
	}

	var p Point
	var tile Tile
	var i bool

	for _, test := range tests {
		tile = Tile{Z: test.z, X: test.x, Y: test.y}
		p = Point{Lat: test.lat, Lng: test.lng}
		i, _ = p.InTile(tile)

		if !i {
			t.Errorf("Expected: %v, got: %v", test.expected, i)
		}
	}
}
