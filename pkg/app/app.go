package app

import (
	"context"
	"errors"
	"fmt"
	"io"
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

type App struct {
	l       *logr.Logger
	cfg     *config.Config
	WG      *wg.Exec
	WGQuick *wgquick.Exec
}

func New(l *logr.Logger, cfg *config.Config) *App {
	return &App{
		l:       l,
		cfg:     cfg,
		WG:      wg.New(cfg.WgExecutable),
		WGQuick: wgquick.New(cfg.WgQuickExecutable),
	}
}

// Run starts application
func (app *App) Run(ctx context.Context, lis net.Listener) error {
	var opts []grpc.ServerOption

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServer := grpc.NewServer(opts...)
	grpcService := server.New(app.WG, app.WGQuick, app.l)
	if !grpcService.WG.Verify(app.l) {
		return errors.New("wg executable is broken")
	}
	if !grpcService.WGQuick.Verify(app.l) {
		return errors.New("wg-quick executable is broken")
	}

	wireguardv1.RegisterWireGuardServiceServer(grpcServer, grpcService)
	reflection.Register(grpcServer)
	app.l.V(0).Info("GRPC listen", "port", app.cfg.Port)
	go func() {
		<-ctx.Done()
		cancel()
		grpcServer.GracefulStop()
	}()
	return grpcServer.Serve(lis)
}

// RunDiag runs application diagnostics
func (app *App) RunDiag(out io.Writer) error {
	wq := app.WGQuick
	wqOk := wq.Verify(app.l)
	wqCmd, wqCmdArgs := wq.GetCmd()
	wqFullCmd := []string{wqCmd}
	wqFullCmd = append(wqFullCmd, wqCmdArgs...)

	wg := app.WG
	wgVersion, wgErr := wg.Version(app.l)
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
