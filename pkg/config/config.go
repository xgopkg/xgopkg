package config

import (
	"github.com/spf13/viper"
	"gopkg.in/logger.v1"
)

func init() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("xgp")
	viper.BindEnv("config_dir")
	viper.BindEnv("env")
	configDir := viper.GetString("config_dir")

	configName := "config"
	if env := viper.GetString("env"); env != "" {
		configName = configName + "." + env
	}
	viper.SetConfigName(configName)

	viper.AddConfigPath(configDir)
	//todo multipconfig path
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../../conf")
	viper.AddConfigPath("../../conf")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("/etc/xgopkg/conf.d")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config error ", err)
	}
	log.Infof("Load config from file:[%s]", viper.ConfigFileUsed())
}

// MySQLURL provides db address for connecting db
func MySQLURL() string {
	return viper.GetString("database.url")
}

// LoggerLevel get looger level from config file
func LoggerLevel() int {
	return viper.GetInt("logger.level")
}

var srvcfg *ServerCfg

//ServerConfig server config
func ServerConfig() *ServerCfg {
	if srvcfg == nil {
		srvcfg = &ServerCfg{}
		srvcfg.APIServer.Port = viper.GetString("server.apiserver.port")
		srvcfg.Engine.Port = viper.GetString("server.engine.port")
	}
	return srvcfg
}

//ServerCfg server config
type ServerCfg struct {
	Engine    engine
	APIServer apiServer
}

//Engine engine
type engine struct {
	Port string
}

//APIServer api server config
type apiServer struct {
	Port string
}
