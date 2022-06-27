package app

import (
	"errors"
	"fmt"
	"log"
	"net"

	wireguardv1 "github.com/ezh/wireguard-grpc/api/wireguard/v1"
	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/server"
	"github.com/ezh/wireguard-grpc/pkg/wg"
	wgquick "github.com/ezh/wireguard-grpc/pkg/wg-quick"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run creates objects via constructors.
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
