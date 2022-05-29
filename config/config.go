package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/cobra"
)

type Config struct {
	WgExecutable string `yaml:"wg_exe" env:"WG_EXE" env-default:"wg" env-description:"wireguard executable file location"`
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
			fmt.Println(help)
		},
	}
}
