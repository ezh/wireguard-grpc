package main

import (
	"fmt"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Run: func(*cobra.Command, []string) {
		help, _ := cleanenv.GetDescription(&config.Config{}, nil)
		fmt.Println(help) // nolint:forbidigo
	},
	Short: "Prints environment variables.",
	Use:   "env",
}
