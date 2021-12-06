package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/ldmtam/basic-go-plugin-example/shared"
)

func main() {
	files, err := ioutil.ReadDir("./plugins")
	if err != nil {
		panic(err)
	}

	var fileName string

	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Error,
	})

	pluginMap := map[string]plugin.Plugin{}
	for _, file := range files {
		fileName = file.Name()
		pluginMap[fileName] = &shared.GreeterPlugin{}
	}

	clientMap := map[string]*plugin.Client{}
	for _, file := range files {
		fileName = file.Name()
		clientMap[fileName] = plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: shared.HandshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command("./plugins/" + fileName),
			Logger:          logger,
		})
	}

	for name, client := range clientMap {
		defer client.Kill()

		rpcClient, err := client.Client()
		if err != nil {
			log.Fatal(err)
		}

		raw, err := rpcClient.Dispense(name)
		if err != nil {
			log.Fatal(err)
		}

		greeter := raw.(shared.Greeter)
		fmt.Printf("Plugin `%v`: %v\n", name, greeter.Greet())
	}
}
