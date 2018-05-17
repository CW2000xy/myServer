package main

import(
	"flag"
	"lib/public_config"
	"server_console/conf"
	"lib/common"
	"github.com/astaxie/beego"
)

func InitLog()  {
	common.SetLogOption(conf.CenterCfg.CfgLog.LogDir,conf.CenterCfg.CfgLog.LogFileName,
		conf.CenterCfg.CfgLog.LogLevel,
		conf.CenterCfg.CfgLog.LogToConsole)

	beego.Info("centerServer log-options set done!")
}

func main()  {
	var pubConfPath string
	flag.StringVar(&pubConfPath, "public_conf", "./public_config.toml", "public_config path")
	flag.Parse()

	public_config.ParseToml(pubConfPath)
	conf.ReadConf("conf/center.toml")
	InitLog()

	db.InitRedis(public_config.GetCfgRedis().Address, public_config.GetCfgRedis().Password)

}