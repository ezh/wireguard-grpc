package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/internal/l"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/hashicorp/go-sockaddr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	RunE:          serverRunE,
	Short:         "Run GRPC server.",
	Use:           "server",
	SilenceErrors: true,
}

func serverRunE(cmd *cobra.Command, args []string) error {
	cfg := config.ReadConfig()

	flags, err := parsePersistentFlags(cmd, cfg)
	if err != nil {
		return err
	}
	app.RegisterLogger(flags.l)

	if flags.wgCmd != "" {
		cfg.WgExecutable = flags.wgCmd
	}
	if flags.wqCmd != "" {
		cfg.WgQuickExecutable = flags.wqCmd
	}
	sa, err := sockaddr.NewSockAddr(cfg.Listen)
	if err != nil {
		log.Fatalf("error parsing listen address: %v", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	var lis net.Listener
	if sa.Type() == sockaddr.TypeUnknown {
		return errors.Wrapf(err, "unable to process listen address: %v", cfg.Listen)
	}
	lis, err = net.Listen(sa.ListenStreamArgs())
	if err != nil {
		return errors.Wrapf(err, "failed to listen address: %v", cfg.Listen)
	}
	if ul, ok := lis.(*net.UnixListener); ok {
		ul.SetUnlinkOnClose(true)
	}
	defer lis.Close()

	// UNIX Socket: Goroutine for checking system signals
	go func() {
		for sig := range c {
			l.Info("signal detected, cleaning up unix", "signal", sig)
			if err := lis.Close(); err != nil {
				l.Error(err, "unix dirty close")
			}
		}
	}()

	return app.New(cfg).RunServer(context.Background(), lis)
}
