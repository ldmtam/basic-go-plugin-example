package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/ldmtam/basic-go-plugin-example/shared"
)

var pluginName = "hi"

type GreeterHi struct {
	logger hclog.Logger
}

func (g *GreeterHi) Greet() string {
	return "Hello con kẹc"
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Error,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	hi := &GreeterHi{logger: logger}

	pluginMap := map[string]plugin.Plugin{
		pluginName: &shared.GreeterPlugin{Impl: hi},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
