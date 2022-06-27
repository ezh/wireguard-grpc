package main

import (
	"fmt"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/cobra"
)

var helpFn = help(rootCmd, rootCmd.HelpFunc())

func help(rootCmd *cobra.Command, orig func(*cobra.Command, []string)) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, output []string) {
		orig(cmd, output)
		if cmd == rootCmd {
			fmt.Fprintln(cmd.OutOrStdout())
			help, _ := cleanenv.GetDescription(&config.Config{}, nil)
			fmt.Fprintln(cmd.OutOrStdout(), help)
		}
	}
}
