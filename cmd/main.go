package main

import (
	"fmt"
	"log"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/internal/logger/zap"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/cobra"
)

var appname = "wireguard-grpc"

func main() {
	rootCmd := &cobra.Command{
		RunE:         run,
		Short:        fmt.Sprintf("%s is a wireguard GRPC API", appname),
		SilenceUsage: true,
	}
	envCmd := config.NewConfigEnvCommand(config.Config{})

	rootCmd.AddCommand(envCmd)
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "silences all output; takes precedence over any verbose setting")
	rootCmd.PersistentFlags().BoolP("superuser", "s", false, "use sudo for wireguard executables")
	rootCmd.PersistentFlags().CountP("verbose", "v", "verbosity. Set this flag multiple times for more verbosity")
	rootCmd.PersistentFlags().StringP("wireguard", "w", "wg", "wireguard executable file")
	rootCmd.SetHelpFunc(help(rootCmd, rootCmd.HelpFunc()))

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
	sudo, err := cmd.Flags().GetBool("superuser")
	if err != nil {
		return err
	}
	if sudo {
		cfg.Sudo = true
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

func help(rootCmd *cobra.Command, orig func(*cobra.Command, []string)) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, output []string) {
		orig(cmd, output)
		if cmd == rootCmd {
			cmd.PrintErrln("")
			help, _ := cleanenv.GetDescription(&config.Config{}, nil)
			cmd.PrintErrln(help)
		}
	}
}
