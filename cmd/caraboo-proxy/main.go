package main

import (
	"fmt"
	"os"

	goConfig "github.com/liampulles/go-config"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/wire"
)

func main() {
	cfg, err := wire.LoadConfig(goConfig.NewEnvSource())
	if err != nil {
		logErr(err)
		os.Exit(1)
	}
	app := wire.Wire(cfg)

	if err := app.Run(); err != nil {
		logErr(err)
		os.Exit(2)
	}
}

func logErr(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
}
