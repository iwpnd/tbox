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

  fmt.Printf("%+v\n", tile)
}
```

Results in

```
{Z:15 X:17600 Y:10786}
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
  bbox := tile.ToBox()

  fmt.Printf("%+v\n", bbox)
}
```

```
{MinLng:13.359375 MinLat:52.24798298528185 MaxLng:13.370361328125 MaxLat:52.25470880113082}
```

### func (t Tile) ContainsPoint(p Point)

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

  ok, err := tile.ContainsPoint(p)

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

### func (t Tile) ToPoint()

Returns center point of a given tile.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/tbox"
  )

func main() {
  tile := tbox.Tile{Z:15, X: 17600, Y: 10786}

  p := tile.ToPoint(p)

  fmt.Printf("%+v\n", p)
}
```

Returns

```
{Lng:13.3648681640625 Lat:52.251345893206334}
```

## License

MIT

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/tbox](https://github.com/iwpnd/tbox)
