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
