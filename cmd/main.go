package main

import (
	"log"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/internal/logger/zap"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/spf13/cobra"
)

var appname = "wireguard-grpc"

func main() {
	rootCmd := &cobra.Command{
		Short: appname,
		RunE:  run,
	}
	envCommand := config.NewConfigEnvCommand(config.Config{})

	rootCmd.AddCommand(envCommand)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("run")
}

func run(cmd *cobra.Command, args []string) error {
	app.Run(zap.New, config.ReadConfig())
	return nil
}
