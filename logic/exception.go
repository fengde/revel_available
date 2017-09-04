package logic

import (
	"revel_available/global"
	"runtime"
)

const LINE_NUM = 32

func CatchException() {
	if err := recover(); err != nil {
		global.Logger.Error("%v", err)
		for i := 0; i < LINE_NUM; i++ {
			funcName, file, line, ok := runtime.Caller(i)
			if ok {
				global.Logger.Error("frame %v:[func:%v,file:%v,line:%v]",
					i,
					runtime.FuncForPC(funcName).Name(),
					file,
					line,
				)
			}
		}
	}
}
