package main

import (
	"github.com/artem-xox/go-shorty/internal/server"
)

func main() {
	go server.StartHTTPServer()
	select {}
}
