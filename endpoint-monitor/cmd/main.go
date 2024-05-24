package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	opservice "github.com/tokamak-network/tokamak-thanos/op-service"
	oplog "github.com/tokamak-network/tokamak-thanos/op-service/log"
	"github.com/urfave/cli/v2"

	endpointMonitor "github.com/tokamak-network/tokamak-thanos/endpoint-monitor"
)

var (
	Version   = ""
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = endpointMonitor.CLIFlags("ENDPOINT_MONITOR")
	app.Version = opservice.FormatVersion(Version, GitCommit, GitDate, "")
	app.Name = "endpoint-monitor"
	app.Usage = "Endpoint Monitoring Service"
	app.Description = ""

	app.Action = endpointMonitor.Main(Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
