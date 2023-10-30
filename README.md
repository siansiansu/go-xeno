# Xeno-Canto CLI

[![GoDoc](https://godoc.org/github.com/siansiansu/go-xeno?status.svg)](http://godoc.org/github.com/siansiansu/go-xeno) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`go-xeno` is a cli tool for integrating [Xeno-Canto](https://xeno-canto.org/) API V2.

## Installation

Make sure you have Go installed. Then, run the following command:

```bash
go get -u github.com/siansiansu/go-xeno/cmd/xeno
```

This command will download the xeno tool to your Go workspace.

Use homebrew:

```bash
brew tap siansiansu/xeno
brew install xeno
```

## Usage

To use `xeno`, run the following command:

Download all the audio recordings.

```bash
xeno "Eurasian Tree Sparrow"
```

Only download 1 audio recording.

```bash
xeno "Eurasian Tree Sparrow" --max-results 1
```

View help and available commands:

```bash
xeno help
```

## Package

Here's an example demonstrating the usage of the `go-xeno` package in Go code:

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
