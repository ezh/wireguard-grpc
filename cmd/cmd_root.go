package main

import (
	"fmt"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	RunE:         rootRunE,
	Short:        fmt.Sprintf("%s is a wireguard GRPC API", appname),
	SilenceUsage: true,
}

func rootRunE(cmd *cobra.Command, args []string) error {
	cfg := config.ReadConfig()

	flags, err := parsePersistentFlags(cmd, cfg)
	if err != nil {
		return err
	}
	if flags.wgCmd != "" {
		cfg.WgExecutable = flags.wgCmd
	}
	if flags.wqCmd != "" {
		cfg.WgQuickExecutable = flags.wqCmd
	}

	return app.Run(flags.l, cfg)
}
