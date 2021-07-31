package tbox

// Tile ...
type Tile struct {
	Z, X, Y int
}

// Tilebox ...
type Tilebox struct {
	MinLng, MinLat, MaxLng, MaxLat float64
}

// NewTile creates a new Tile
func NewTile(z, x, y int) *Tile {
	return &Tile{z, x, y}
}

// ToBox returns the bounding box of a given Tile
func (t Tile) ToBox() *Tilebox {
	return &Tilebox{
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
