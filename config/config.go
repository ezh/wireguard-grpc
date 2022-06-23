package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/cobra"
)

type Config struct {
	LogLevel          string `yaml:"loglevel" env:"LOG_LEVEL" env-description:"application log level"`
	Port              int    `yaml:"port" env:"PORT" env-default:"8081" env-description:"GRPC server port"`
	WgExecutable      string `yaml:"wg" env:"WG_EXE" env-default:"wg" env-description:"wireguard executable"`
	WgQuickExecutable string `yaml:"wgquick" env:"WGQUICK_EXE" env-default:"wg-quick" env-description:"wg-quick wrapper"`
	Sudo              bool   `yaml:"sudo" env:"SUDO" env-defalut:"false" env-description:"use sudo for executables"`
}

// ReadConfig reads configuration
func ReadConfig() *Config {
	cfg := Config{}
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Unable to read configuration %v", err)
	}
	return &cfg
}

func NewConfigEnvCommand(c Config) *cobra.Command {
	return &cobra.Command{
		Use:   "env",
		Short: "Prints environment variables.",
		Run: func(*cobra.Command, []string) {
			help, _ := cleanenv.GetDescription(&Config{}, nil)
			fmt.Println(help) // nolint:forbidigo
		},
	}
}
