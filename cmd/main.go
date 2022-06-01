package main

import (
	"fmt"
	"log"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/internal/logger/zap"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/spf13/cobra"
)

var appname = "wireguard-grpc"

func main() {
	rootCmd := &cobra.Command{
		Short: fmt.Sprintf("%s is a wireguard GRPC API", appname),
		RunE:  run,
	}
	envCmd := config.NewConfigEnvCommand(config.Config{})

	rootCmd.AddCommand(envCmd)
	rootCmd.PersistentFlags().StringP("wireguard", "w", "wg", "wireguard executable file")
	rootCmd.PersistentFlags().CountP("verbose", "v", "verbosity. Set this flag multiple times for more verbosity")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Silences all output; takes precedence over any verbose setting")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	cfg := config.ReadConfig()
	verbosity, err := cmd.Flags().GetCount("verbose")
	if err != nil {
		return err
	}
	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return err
	}
	wgExe, err := cmd.Flags().GetString("wireguard")
	if err != nil {
		return err
	}
	// replace WgExecutable if user set -wg
	if wgExe != "wg" {
		cfg.WgExecutable = wgExe
	}
	if quiet {
		return app.Run(zap.New, cfg, -1)
	}
	return app.Run(zap.New, cfg, verbosity)
}
