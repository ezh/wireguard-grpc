package main

import (
	"errors"
	"fmt"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/wg"
	wgquick "github.com/ezh/wireguard-grpc/pkg/wg-quick"
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

	wq := wgquick.New(cfg.WgQuickExecutable)
	wqOk := wq.Verify(flags.l)
	wqCmd := []string{wq.Cmd}
	wqCmd = append(wqCmd, wq.CmdArgs...)

	wg := wg.New(cfg.WgExecutable)
	wgVersion, wgErr := wg.Version(flags.l)
	wgCmd := []string{wg.Cmd}
	wgCmd = append(wgCmd, wg.CmdArgs...)

	fmt.Fprintf(cmd.OutOrStdout(),
		"wg correct: %v\nwg version: %s\nwg cmd: %v\n\n"+
			"wg-quick correct: %v\nwg-quick cmd: %v\n",
		len(wgVersion) > 0, wgVersion, wgCmd, wqOk, wqCmd)

	if wgErr != nil {
		return wgErr
	}
	if !wqOk {
		return errors.New("wg-quick is broken")
	}
	return nil
}
