package main

import (
	"log"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/internal/logger/zap"
	"github.com/ezh/wireguard-grpc/pkg/logger"
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
)

var appname = "wireguard-grpc"

// logBuilder is hard coded type of logger implementation
var logBuilder = zap.New

// newLogger creates a new logger
func newLogger(rawLogLevel string, verbosity int) (logr.Logger, error) {
	return logger.NewLogger(logBuilder, rawLogLevel, verbosity)
}

type persistentFlags struct {
	wgCmd string
	wqCmd string
	l     *logr.Logger
}

func parsePersistentFlags(cmd *cobra.Command, cfg *config.Config) (*persistentFlags, error) {
	f := persistentFlags{}

	verbosity, err := cmd.Flags().GetCount("verbose")
	if err != nil {
		return nil, err
	}
	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return nil, err
	}
	wgExe, err := cmd.Flags().GetString("wireguard")
	if err != nil {
		return nil, err
	}
	// replace WgExecutable if user set -wg
	if wgExe != "wg" {
		f.wgCmd = wgExe
	}

	var l logr.Logger
	if quiet {
		l, err = newLogger(cfg.LogLevel, -1)
	} else {
		l, err = newLogger(cfg.LogLevel, verbosity)
	}
	if err != nil {
		return nil, err
	}
	f.l = &l

	return &f, nil
}

func main() {
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(diagCmd)
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "silences all output; takes precedence over any verbose setting")
	rootCmd.PersistentFlags().CountP("verbose", "v", "verbosity. Set this flag multiple times for more verbosity")
	rootCmd.PersistentFlags().StringP("wireguard", "w", "wg", "wireguard executable file")
	rootCmd.SetHelpFunc(helpFn)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
