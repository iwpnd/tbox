# tbox

Tbox provides methods to interact with map tiles described by Z/X/Y coordinates.

- get bounding box for tile
- get tiles from a bounding box
- get parent tile for a tile
- get children tiles for a tile
- get tile coordinates for a given point at a given zoom level
- get the center point of a given tile
- validate if a given point is within a given tile
- validate if a given tile contains a given point

## installation

```
go get -u github.com/iwpnd/tbox
```

## usage

### func (p Point) ToTile(zoom int)

Get the Tile coordinates for a given a point at a given zoom level.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  p := tbox.Point{Lat: 52.25, Lng: 13.37}
  tile, err := p.ToTile(15)

  if err != nil {
      fmt.Println(err)
    }

  fmt.Printf("%+v\n", tile)
}
```

Results in

```
{Z:15 X:17600 Y:10786}
```

### func (p Point) Intersects(tile Tile)

Check whether a given point intersects a given tile defined by Z/X/Y coordinates.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  p := tbox.Point{Lat: 52.25, Lng: 13.37}
  tile := tbox.Tile{Z:15, X: 17600, Y: 10786}

  ok, err := p.Intersects(tile)

  if err != nil {
      fmt.Println(err)
    }

  fmt.Println(ok)
}
```

Results in

```
true
```

### func (t Tile) Bbox()

Get the bounding box of a given tile defined by Z/X/Y.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  tile := tbox.Tile{Z:15, X: 17600, Y: 10786}
  bbox := tile.ToBox()

  fmt.Printf("%+v\n", bbox)
}
```

```
{MinLng:13.359375 MinLat:52.24798298528185 MaxLng:13.370361328125 MaxLat:52.25470880113082}
```

### func (t Tile) Contains(p Point)

Check if Point is in tile.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  tile := tbox.Tile{Z:15, X: 17600, Y: 10786}
  p := tbox.Point{Lat: 52.25, Lng: 13.37}

  ok, err := tile.Contains(p)

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(ok)
}
```

Returns

```
true
```

### func (t Tile) Center()

Returns center point of a given tile.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  tile := tbox.Tile{Z:15, X: 17600, Y: 10786}

  p := tile.Center(p)

  fmt.Printf("%+v\n", p)
}
```

Returns

```
{Lng:13.3648681640625 Lat:52.251345893206334}
```

### func (t Tile) Children() 

Returns an array of child tiles.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  tile := tbox.Tile{Z:0, X: 0, Y: 0}
  children := tile.Children()

  fmt.Printf("%+v\n", p)
}
```

Returns

```
[{Z:1,X:0,Y:0},{Z:1,X:1,Y:0},{Z:1,X:0,Y:1},{Z:1,X:1,Y:1}]
```

### func (t Tile) Parent() 

Returns a the parent tile. If zoom level 0, returns tile.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  tile := tbox.Tile{Z:1, X: 0, Y: 0}
  parent := tile.Parent()

  fmt.Printf("%+v\n", p)
}
```

Returns

```
{Z:0,X:0,Y:0}
```

### func FromBounds(b BoundingBox)

Returns all tiles intersecting the input bounding box at a certain zoom level

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  bbox := tbox.BoundingBox{MinLng: 10.045, MinLat: 51.2114, MaxLng: 13.825, MaxLat: 53.575}
  tiles := tbox.FromBounds(bbox, 7)

  fmt.Printf("%+v\n", tiles)
}
```

Returns

```
[
  {X: 67, Y: 41, Z: 7},
	{X: 68, Y: 41, Z: 7},
	{X: 67, Y: 42, Z: 7},
	{X: 68, Y: 42, Z: 7},
]
```



## License

MIT

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/tbox](https://github.com/iwpnd/tbox)
