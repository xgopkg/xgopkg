package main

import (
	"net/http"

	_ "github.com/go-xorm/xorm"
	_ "github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/mysql"
	_ "github.com/spf13/viper"
	"gopkg.in/logger.v1"

	"xgopkg.com/xgopkg/pkg/config"
	"xgopkg.com/xgopkg/pkg/mapper"
	"xgopkg.com/xgopkg/pkg/migrate"
	"xgopkg.com/xgopkg/pkg/route"
)

func init() {

}
func main() {
	//set log level
	log.SetOutputLevel(config.LoggerLevel())
	//db connection
	mapper.Connect()
	defer mapper.Close()
	//db migrate
	migrate.Migrate()

	//server addr
	var engineAddr, apiserverAddr string
	//get server config
	serverCfg := config.ServerConfig()

	//run apiserver
	go func() {
		if serverCfg.APIServer.Port != "" {
			apiserverAddr = ":" + serverCfg.APIServer.Port
		} else {
			log.Fatal("apiserver port not configured.")
		}
		route.API()
		route.WebUI()
		log.Fatal(http.ListenAndServe(apiserverAddr, route.APIServerContainer()))
	}()

	if serverCfg.Engine.Port != "" {
		engineAddr = ":" + serverCfg.Engine.Port
	} else {
		log.Fatal("Engine server port not configured.")
	}

	route.Package()
	//Run main server
	log.Fatal(http.ListenAndServe(engineAddr, nil))
}
