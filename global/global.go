package global

import (
	"fmt"

	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego/logs"
)

var (
	// 配置文件路径
	ConfigFile = "revel_available/conf/app.conf"
	// 配置对象
	Config *goconfig.ConfigFile
	// 日志对象
	Logger = logs.NewLogger()
)

func GlobalInit() {
	Config, _ = goconfig.LoadConfigFile(ConfigFile)
	fileSwitch := Config.MustValue("log", "fileswitch")
	level := Config.MustValue("log", "level")
	if fileSwitch == "on" {
		fileName := Config.MustValue("log", "filename")
		daily := Config.MustValue("log", "daily")
		maxdays := Config.MustValue("log", "maxdays")
		Logger.SetLogger(logs.AdapterFile,
			fmt.Sprintf(`{"filename": "%v", "daily": %v, "maxdays": %v, "level": %v}`, fileName, daily, maxdays, level))
	} else {
		Logger.SetLogger(logs.AdapterConsole)
	}
	// 设置日志输出文件名和文件行号
	Logger.EnableFuncCallDepth(true)
	// 设置日志是否异步输出
	if Config.MustValue("log", "async") == "true" {
		// 设置异步输出日志，并且设置缓存大小1e3
		Logger.Async(1e3)
	}

	Logger.Info("日志系统正常启动")
}
