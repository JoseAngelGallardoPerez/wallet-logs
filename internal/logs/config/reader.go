package config

import (
	"github.com/Confialink/wallet-pkg-env_config"
	"github.com/Confialink/wallet-pkg-env_mods"
)

// readConfig reads configs from ENV variables
func ReadConfig() *Config {
	defaultConfigReader := env_config.NewReader("logs")
	return &Config{
		Port:         env_config.Env("VELMIE_WALLET_LOGS_PORT", ""),
		ProtobufPort: env_config.Env("VELMIE_WALLET_LOGS_PROTOBUF_PORT", ""),
		Env:          env_config.Env("ENV", env_mods.Development),
		Cors:         defaultConfigReader.ReadCorsConfig(),
		Db:           defaultConfigReader.ReadDbConfig(),
	}
}
