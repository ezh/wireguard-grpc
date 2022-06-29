package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	wireguardv1 "github.com/ezh/wireguard-grpc/api/wireguard/v1"
	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/server"
	"github.com/ezh/wireguard-grpc/pkg/wg"
	wgquick "github.com/ezh/wireguard-grpc/pkg/wg-quick"
	"github.com/go-logr/logr"
	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run starts application
func Run(l *logr.Logger, cfg *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpcService := server.GRPC{
		WG:      wg.New(cfg.WgExecutable),
		WGQuick: wgquick.New(cfg.WgQuickExecutable),
	}
	if !grpcService.WG.Verify(l) {
		return errors.New("wg executable is broken")
	}
	if !grpcService.WGQuick.Verify(l) {
		return errors.New("wg-quick executable is broken")
	}

	wireguardv1.RegisterWireGuardServiceServer(grpcServer, &grpcService)
	reflection.Register(grpcServer)
	l.V(0).Info("GRPC listen", "port", cfg.Port)
	return grpcServer.Serve(lis)
}

// RunDiag runs application diagnostics
func RunDiag(l *logr.Logger, cfg *config.Config, out io.Writer) error {
	wq := wgquick.New(cfg.WgQuickExecutable)
	wqOk := wq.Verify(l)
	wqCmd, wqCmdArgs := wq.GetCmd()
	wqFullCmd := []string{wqCmd}
	wqFullCmd = append(wqFullCmd, wqCmdArgs...)

	wg := wg.New(cfg.WgExecutable)
	wgVersion, wgErr := wg.Version(l)
	wgCmd, wgCmdArgs := wg.GetCmd()
	wgFullCmd := []string{wgCmd}
	wgFullCmd = append(wgFullCmd, wgCmdArgs...)

	fmt.Fprintf(out,
		"wg correct: %v\nwg version: %s\nwg cmd: %v\n\n"+
			"wg-quick correct: %v\nwg-quick cmd: %v\n",
		len(wgVersion) > 0, wgVersion, wgFullCmd, wqOk, wqFullCmd)

	if wgErr != nil {
		return wgErr
	}
	if !wqOk {
		return errors.New("wg-quick is broken")
	}
	return nil
}

// RunEnv shows environment variable and actual connfiguration
func RunEnv(cfg *config.Config, out io.Writer) error {
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return err
	}
	help, err := cleanenv.GetDescription(cfg, nil)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, help)
	fmt.Fprintf(out, "Actual configuration: %#v", cfg)
	return nil
}
