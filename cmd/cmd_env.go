package main

import (
	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.RunEnv(config.ReadConfig(), cmd.OutOrStdout())
	},
	Short:         "Prints environment variables.",
	Use:           "env",
	SilenceErrors: true,
}
