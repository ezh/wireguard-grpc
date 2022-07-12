package server

import (
	"context"

	wireguardv1 "github.com/ezh/wireguard-grpc/api/wireguard/v1"
	"github.com/ezh/wireguard-grpc/pkg/utilities"
	"github.com/ezh/wireguard-grpc/pkg/wg"
	wgquick "github.com/ezh/wireguard-grpc/pkg/wg-quick"
	"github.com/go-logr/logr"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPC struct {
	WG      *wg.Exec
	WGQuick *wgquick.Exec
	l       *logr.Logger
	wireguardv1.UnimplementedWireGuardServiceServer
}

var _ wireguardv1.WireGuardServiceServer = (*GRPC)(nil)

func New(wg *wg.Exec, wq *wgquick.Exec, l *logr.Logger) *GRPC {
	return &GRPC{
		WG:      wg,
		WGQuick: wq,
		l:       l,
	}
}

func (*GRPC) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *GRPC) Dump(ctx context.Context, req *wireguardv1.DumpRequest) (*wireguardv1.DumpResponse, error) {
	var response wireguardv1.DumpResponse

	names := req.GetWgIfNames()
	devices, err := s.WG.Dump(s.l)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to dump, %s", err.Error())
	}
	for _, d := range devices {
		if len(names) > 0 && !slices.Contains(names, d.Name) {
			continue
		}
		var peers []*wireguardv1.PeerConfig
		for _, p := range d.Peers {
			peer := wireguardv1.PeerConfig{
				PublicKey:           p.PublicKey[:],
				PresharedKey:        p.PresharedKey[:],
				EndpointIp:          p.Endpoint.IP.String(),
				EndpontPort:         uint32(p.Endpoint.Port),
				PersistentKeepalive: uint32(p.PersistentKeepaliveInterval.Seconds()),
				LastHandshakeTime:   p.LastHandshakeTime.Unix(),
				ReceiveBytes:        p.ReceiveBytes,
				TransmitBytes:       p.TransmitBytes,
				AllowedIps:          utilities.IPNetSliceToString(p.AllowedIPs),
			}
			peers = append(peers, &peer)
		}
		responseInterface := wireguardv1.InterfaceConfig{
			WgIfName:     d.Name,
			PublicKey:    d.PublicKey[:],
			ListenPort:   uint32(d.ListenPort),
			FirewallMark: uint32(d.FirewallMark),
			Peers:        peers,
		}
		response.Interfaces = append(response.Interfaces, &responseInterface)
	}
	return &response, nil
}
