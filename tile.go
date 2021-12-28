package tbox

// Tile ...
type Tile struct {
	Z, X, Y int
}

// BoundingBox ...
type BoundingBox struct {
	MinLng, MinLat, MaxLng, MaxLat float64
}

// ToBox returns the bounding box of a given Tile
func (t Tile) ToBox() *BoundingBox {
	return &BoundingBox{
		MaxLat: tileToLat(t.Y, t.Z),
		MinLng: tileToLng(t.X, t.Z),
		MinLat: tileToLat(t.Y+1, t.Z),
		MaxLng: tileToLng(t.X+1, t.Z),
	}
}

// ContainsPoint valides whether a given Point is within a given Tile
func (t Tile) ContainsPoint(p Point) (bool, error) {
	if !p.Valid() {
		return false, &invalidPointError{p: p, msg: "invalid point"}
	}
	tbox := t.ToBox()

	return p.Lat > tbox.MinLat && p.Lat < tbox.MaxLat && p.Lng > tbox.MinLng && p.Lng < tbox.MaxLng, nil
}

// ToPoint returns the centerpoint of a given Tile
func (t Tile) ToPoint() *Point {
	tbox := t.ToBox()

	cLng := tbox.MinLng + (tbox.MaxLng-tbox.MinLng)/2
	cLat := tbox.MinLat + (tbox.MaxLat-tbox.MinLat)/2

	return &Point{Lng: cLng, Lat: cLat}
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
