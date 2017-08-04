package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"

	ethUtils "github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/params"

	"github.com/tendermint/ethermint/cmd/utils"
	"github.com/tendermint/ethermint/version"
)

var (
	// The app that holds all commands and flags.
	app = ethUtils.NewApp(version.Version, "the ethermint command line interface")

	// flags that configure the ABCI app
	ethermintFlags = []cli.Flag{
		utils.TendermintAddrFlag,
		utils.ABCIAddrFlag,
		utils.ABCIProtocolFlag,
		utils.VerbosityFlag,
		utils.ConfigFileFlag,
	}
)

func init() {
	app.Action = ethermintCmd
	app.HideVersion = true
	app.Commands = []cli.Command{
		{
			Action:      initCmd,
			Name:        "init",
			Usage:       "init genesis.json",
			Description: "Initialize the files",
		},
		{
			Action:      versionCmd,
			Name:        "version",
			Usage:       "",
			Description: "Print the version",
		},
		{
			Action: resetCmd,
			Name:   "unsafe_reset_all",
			Usage:  "(unsafe) Remove ethermint database",
		},
	}

	app.Flags = append(app.Flags, ethermintFlags...)

	app.Before = func(ctx *cli.Context) error {
		if err := utils.Setup(ctx); err != nil {
			return err
		}

		ethUtils.SetupNetwork(ctx)

		return nil
	}
}

func versionCmd(ctx *cli.Context) error {
	fmt.Println("ethermint: ", version.Version)
	fmt.Println("go-ethereum: ", params.Version)
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
