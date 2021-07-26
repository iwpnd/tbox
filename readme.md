# tbox

## installation

```
go get -u github.com/iwpnd/tbox
```

## usage

```go
package main

import (
  "fmt",

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

### Example

## License

MIT

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/tbox](https://github.com/iwpnd/tbox)
