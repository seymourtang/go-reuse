# go-reuse
[![Go Report Card](https://goreportcard.com/badge/github.com/seymourtang/go-reuse)](https://goreportcard.com/report/github.com/seymourtang/go-reuse)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/seymourtang/go-reuse)
[![Go Reference](https://pkg.go.dev/badge/github.com/seymourtang/go-reuse.svg)](https://pkg.go.dev/github.com/seymourtang/go-reuse)
![GitHub](https://img.shields.io/github/license/seymourtang/go-reuse)


Package go-reuse provides Listen functions that set socket options in order to be able to reuse ports on Windows,Linux and macOS platform.

## Installation
```shell
go get github.com/seymourtang/go-reuse
```

## Getting Started

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/seymourtang/go-reuse"
)

func main() {
	listener, err := reuse.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		_ = listener.Close()
	}()

	s := &http.Server{}
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "gid:%d,pid:%d", os.Getegid(), os.Getpid())
	})

	log.Fatalln(s.Serve(listener))
}

```
