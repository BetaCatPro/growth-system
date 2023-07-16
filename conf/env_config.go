package conf

import (
	"encoding/json"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// GlobalConfig project config info
// maybe saved in os.Env or k8s service config yaml file.
var GlobalConfig *ProjectConfig

// envConfigName env config name for project
const envConfigName = "GROWTH_CONFIG"

type Mysql struct {
	Engine          string
	Username        string
	Password        string
	Host            string
	Port            int
	Database        string
	Charset         string
	ShowSql         bool
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

// ProjectConfig project's config
type ProjectConfig struct {
	Db    Mysql
	Cache struct{}
}

// LoadConfigs load global config info
func LoadConfigs() {
	LoadEnvConfig()
}

// LoadEnvConfig load configs from env config with name of envConfigName, json format
func LoadEnvConfig() {
	pc := &ProjectConfig{}
	// load from os env
	if strConfigs := os.Getenv(envConfigName); len(strConfigs) > 0 {
		if err := json.Unmarshal([]byte(strConfigs), pc); err != nil {
			zap.S().Errorf("conf.LoadEnvConfig(%s) error=%s\n", envConfigName, err.Error())
			return
		}
	} else {
		// load from config file
		viper.SetConfigName("mysql")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")

		if err := viper.ReadInConfig(); err != nil {
			zap.S().Errorf("fatal error config file: %w", err)
			return
		}

		if err := viper.Unmarshal(pc); err != nil {
			zap.S().Errorf("conf.LoadYmlConfig error=%s\n", err.Error())
			return
		}
	}

	if pc == nil || pc.Db.Username == "" { // no config info
		zap.S().Errorf("empty os.Getenv config ", envConfigName)
		return
	}

	GlobalConfig = pc
}
