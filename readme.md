# tbox

## installation

```
go get -u github.com/iwpnd/tbox
```

## usage

```
package main

import (
  "fmt",

  "github.com/iwpnd/tbox"
  )

func main() {
  p := tbox.Point{Lat: 52.25, Lng: 13.37}
  tile := p.ToTile(15)

  fmt.Println(tile)
}
```

### Example

## License

MIT

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/tbox](https://github.com/iwpnd/tbox)
