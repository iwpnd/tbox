# tbox

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

  fmt.Println(tile)
}
```

Results in

```
{15 17600 10786} // Z/X/Y
```

### func (p Point) InTile(tile Tile)

Check whether a given point is within a given tile defined by Z/X/Y coordinates.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  p := tbox.Point{Lat: 52.25, Lng: 13.37}
  tile := tbox.Tile{Z:15, X: 17600, Y: 10786}

  ok, err := p.InTile(tile)

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

### func (t Tile) ToBox()

Get the bounding box of a given tile defined by Z/X/Y.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  tile := tbox.Tile{Z:15, X: 17600, Y: 10786}

  bbox, err := tile.ToBox()

  if err != nil {
      fmt.Println(err)
    }

  fmt.Println(bbox)
}
```

## License

MIT

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/tbox](https://github.com/iwpnd/tbox)
