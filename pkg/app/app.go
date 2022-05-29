package app

import (
	"log"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/logger"
)

// Run creates objects via constructors.
func Run(logBuilder logger.LogBuilder, cfg *config.Config) {
	err, l := logBuilder(0)
	if err != nil {
		log.Fatal(err)
	}
	_ = l
}
