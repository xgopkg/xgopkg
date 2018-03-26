package mapper

import (
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"gopkg.in/logger.v1"
)

var engine *xorm.Engine

//Connect DB
func Connect() {
	if engine != nil {
		return
	}
	connectDb()

}
func connectDb() {
	driver := viper.GetString("database.type")
	url := viper.GetString("database.url")
	maxIdle := viper.GetInt("database.maxIdle")
	maxActive := viper.GetInt("database.maxActive")
	log.Debugf("driver: %s, url: %s", driver, url)

	// TODO: Check cofig is nil
	// FIXME:
	var err error
	engine, err = xorm.NewEngine(driver, url)
	if err != nil {
		log.Fatal("DB connect error ", err)
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "xgp_")
	engine.SetTableMapper(tbMapper)
	engine.SetMaxIdleConns(maxIdle)
	engine.SetMaxOpenConns(maxActive)
}
