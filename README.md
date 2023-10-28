# Xeno-Canto Client for Go

[![GoDoc](https://godoc.org/github.com/siansiansu/go-xeno?status.svg)](http://godoc.org/github.com/siansiansu/go-xeno) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go-Xeno is a simple client for integrating Xeno-Canto API V2.

## Installation

To use Go-Xeno in your Go module, you can simply run:

```bash
go get -u github.com/siansiansu/go-xeno
```

## Usage

Here's an example of how you can use Go-Xeno:

```go
package main

import (
  "context"
  "fmt"

  "github.com/siansiansu/go-xeno"
)

func main() {
  var ctx = context.Background()
  client, err := xeno.NewClient()
  if err != nil {
    panic(err)
  }
  r, err := client.Get(ctx, "Taiwan blue magpie", xeno.Page(1), xeno.NumPages(1))
  if err != nil {
    panic(err)
  }
  for _, e := range r.Recordings {
    fmt.Println(e.Rec, e.Loc, e.File)
  }
}
```

## Contributing

Contributions are welcome! Report bugs or request features by opening an issue. If you want to contribute code, fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
