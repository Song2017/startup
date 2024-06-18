package initial

import (
	"encoding/json"
	"os"
	_pkg "startup/_pkg"
	"sync"
)

var (
	mu_server_config sync.Mutex
)

type ServerConfig struct {
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	PgConn PgConn `mapstructure:"pg-conn" json:"pg-conn" yaml:"pg-conn"`

	ServerName string `mapstructure:"server-name" json:"server-name" yaml:"server-name"`
	LogPath    string `mapstructure:"log-path" json:"log-path" yaml:"log-path"`
	TimeOut    int    `mapstructure:"time-out" json:"time-out" yaml:"time-out"`

	SecurityKey   string `mapstructure:"security-key" json:"security-key" yaml:"security-key"`
	SecurityValue string `mapstructure:"security-value" json:"security-value" yaml:"security-value"`
}

func LoadServerConfig(serverConfig *ServerConfig) *ServerConfig {
	if serverConfig == nil {
		mu_server_config.Lock()
		defer mu_server_config.Unlock()
		if serverConfig == nil {
			if conf, err := _pkg.B64Decode(
				os.Getenv("SERVER_CONFIG")); err == nil {
				err = json.Unmarshal([]byte(conf), &serverConfig)
				if err != nil {
					panic("Load Server Config Error: " + err.Error())
				}
			}
			if serverConfig.SecurityKey == "" {
				serverConfig.SecurityKey = "Authentication"
			}
		}
	}

	return serverConfig
}

func InitServerConfig() {
	O_SERVER_CONFIG = LoadServerConfig(O_SERVER_CONFIG)
	O_REDIS = O_SERVER_CONFIG.Redis.InitRedisByConfig()
	O_DB_PG = O_SERVER_CONFIG.PgConn.GormPgSqlByConfig()
}
