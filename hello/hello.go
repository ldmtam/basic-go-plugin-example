package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/ldmtam/basic-go-plugin-example/shared"
)

var pluginName = "hello"

type GreeterHello struct {
	logger hclog.Logger
}

func (g *GreeterHello) Greet() string {
	return "Thích chào không thằng cờ hó?"
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Error,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	hello := &GreeterHello{logger: logger}

	pluginMap := map[string]plugin.Plugin{
		pluginName: &shared.GreeterPlugin{Impl: hello},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
