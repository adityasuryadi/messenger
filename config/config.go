package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

type Config struct {
	Service Service `mapstructure:"service"`
	JWT     JWT     `mapstructure:"jwt"`
}

type Service struct {
	Port string `mapstructure:"port"`
}

type JWT struct {
	SecretJWT string `mapstructure:"secretJWT"`
	Ttl       int    `mapstructure:"ttl"`
}

type option struct {
	configFolders []string
	configFile    string
	configType    string
}

type Option func(*option)

var config *Config

func Init(opts ...Option) error {

	opt := &option{
		configFolders: getdefaultConfigFolder(),
		configFile:    getdefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	for _, configFolder := range opt.configFolders {
		viper.AddConfigPath(configFolder)
	}
	// viper.AutomaticEnv()

	config = new(Config)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		slog.Error("failed to unmarshal config", err)
		return err
	}
	return nil
}

func getdefaultConfigFolder() []string {
	return []string{"./config"}
}

func getdefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func SetConfigFolder(configFolders []string) Option {
	return func(o *option) {
		o.configFolders = configFolders
	}
}

func SetConfigFile(configFile string) Option {
	return func(opt *option) {
		opt.configFile = configFile
	}
}

func SetConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
