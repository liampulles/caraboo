package main

import (
	"os"

	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/wire"
)

func main() {
	os.Exit(wire.Run())
}
