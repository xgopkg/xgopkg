package config

import (
	"path"
	"runtime"

	"github.com/spf13/viper"
	"gopkg.in/logger.v1"
)

func init() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("xgp")
	viper.BindEnv("config_dir")
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}
	log.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	configPath := path.Join(path.Dir(filename), "../../conf")
	viper.SetDefault("config_dir", configPath)
	configDir := viper.GetString("config_dir")
	if configDir == "" {
		configDir = "./conf"
	}
	log.Info(configDir)
	viper.SetConfigName("config")
	viper.AddConfigPath(configDir)
	//todo multipconfig path
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config error ", err)
	}
	log.Info(viper.GetString("mysql.host"))
}

//GetConfig get config by viper
// func GetConfig() {
// 	configDir := viper.GetString("config_dir")
// 	if configDir == "" {
// 		configDir = "./conf"
// 	}
// 	log.Info(configDir)
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath(configDir)
// 	//todo multipconfig path
// 	viper.AddConfigPath(".")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatal("read config error ", err)
// 	}
// 	log.Info(viper.GetString("mysql.host"))
// }
