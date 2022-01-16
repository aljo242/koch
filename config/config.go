package config

import (
	"errors"
	"fmt"

	"github.com/aljo242/koch/util/file_util"

	"github.com/spf13/viper"
)

const (
	DefaultConfigPath = "$HOME/.koch/config/"
)

// ErrInvalidConfig indicates that the config file is invalid
var ErrInvalidConfig = errors.New("invalid config")

func New(path string) error {
	// check if path exists
	if !file_util.Exists(path) {
		path = DefaultConfigPath
	}

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()

	// if no config file, write default to default cfg path
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// default config
		err = viper.SafeWriteConfigAs(path) // write current config to path
	}
	if err != nil {
		return fmt.Errorf("fatal error config file %w", err)
	}

	// TODO ensure min config exists
	if !ensureMinConfig() {
		return ErrInvalidConfig
	}

	return nil
}

func ensureMinConfig() bool {

	return true
}

func LogLevel() string {
	return viper.GetString("logger.level")
}

func ServerSecure() bool {
	return viper.GetBool("server.secure")
}

func ServerIP() string {
	return viper.GetString("server.IP")
}

func ServerChooseIP() bool {
	return viper.GetBool("server.chooseIP")
}

func ServerPort() string {
	return viper.GetString("server.port")
}

func ServerCertFile() string {
	return viper.GetString("server.certFile")
}

func ServerKeyFile() string {
	return viper.GetString("server.keyFile")
}

func ServerRootCA() string {
	return viper.GetString("server.rootCA")
}

func ServerHost() string {
	return viper.GetString("server.host")
}

func ServerShutdownCode() int {
	return viper.GetInt("server.shutdownCode")
}

func ServerCmdEnable() bool {
	return viper.GetBool("server.cmdEnable")
}

func ServerCacheMaxAge() int {
	return viper.GetInt("server.cacheMaxAge")
}

func AppName() string {
	return viper.GetString("app")
}

func OwnerName() string {
	return viper.GetString("owner.name")
}

// TODO add more as config gets fleshed out
