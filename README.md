# NanoID

[![Build Status](https://travis-ci.org/bsm/nanoid.png?branch=master)](https://travis-ci.org/bsm/nanoid)
[![GoDoc](https://godoc.org/github.com/bsm/nanoid?status.png)](http://godoc.org/github.com/bsm/nanoid)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsm/nanoid)](https://goreportcard.com/report/github.com/bsm/nanoid)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is a port the original [JavaScript](https://github.com/ai/nanoid) library, a tiny, secure, URL-friendly, unique string ID generator for [Go](https://golang.org).

## Examples

```go
package main

import (
	"fmt"

	"github.com/bsm/nanoid"
)

func main() {
	// Generate a base64-encoded ID with default length.
	fmt.Println(nanoid.New()) // => "AAXTFo7Gdp2kwyvZDySP6"

	// Generate a base64-encoded ID with custom length.
	fmt.Println(nanoid.NewSize(16)) // => "fSYXVjtyrXN8Bw33"

	// Use custom encoding.
	fmt.Println(nanoid.Base32.MustGenerate(16)) // => "WOEJC4M2HXXNSXQ7"
}
```
