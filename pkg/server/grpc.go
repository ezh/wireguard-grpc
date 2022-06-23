package server

import (
	"context"

	wireguardv1 "github.com/ezh/wireguard-grpc/api/wireguard/v1"
	"github.com/ezh/wireguard-grpc/pkg/wg"
	wgquick "github.com/ezh/wireguard-grpc/pkg/wg-quick"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPC struct {
	WG      wg.Exec
	WGQuick wgquick.Exec
	wireguardv1.UnimplementedWireGuardServiceServer
}

var _ wireguardv1.WireGuardServiceServer = (*GRPC)(nil)

func (*GRPC) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
