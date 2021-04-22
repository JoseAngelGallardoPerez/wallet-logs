package config

import (
	"github.com/Confialink/wallet-logs/internal/logs/config/logs"
	"github.com/Confialink/wallet-pkg-env_config"
)

type Config struct {
	Env          string
	Port         string
	ProtobufPort string
	Db           *env_config.Db
	Cors         *env_config.Cors
}

var cfg *Config

func init() {
	cfg = ReadConfig()
	validate(cfg)
}

func GetConfig() *Config {
	return cfg
}

func validate(cfg *Config) {
	logger := logs.Logger.New("service", "configValidator")
	validator := env_config.NewValidator(logger)
	validator.ValidateCors(cfg.Cors, logger)
	validator.ValidateDb(cfg.Db, logger)
	validator.CriticalIfEmpty(cfg.Port, "VELMIE_WALLET_LOGS_PORT", logger)
	validator.CriticalIfEmpty(cfg.ProtobufPort, "VELMIE_WALLET_LOGS_PROTOBUF_PORT", logger)
}
