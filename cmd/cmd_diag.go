package main

import (
	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/spf13/cobra"
)

var diagCmd = &cobra.Command{
	RunE:         diagRunE,
	Short:        "test wireguard-grpc configuration",
	SilenceUsage: true,
	Use:          "diag",
}

func diagRunE(cmd *cobra.Command, args []string) error {
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

	return app.RunDiag(flags.l, cfg, cmd.OutOrStdout())
}
