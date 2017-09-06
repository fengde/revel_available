package controllers

import (
	"errors"
	"net/url"
	"revel_available/define"
	"revel_available/global"
	"strconv"
	"strings"
	"time"

	"github.com/revel/revel"
)

type Base struct {
	*revel.Controller
	In            url.Values
	Cookie        map[string]string
	StartUnixTime int64
	EndUnixTime   int64
	Out           define.Out
}

func (p *Base) After() revel.Result {
	p.EndUnixTime = time.Now().UnixNano()
	global.Logger.Info("返回：%v ", p.Result)
	global.Logger.Info("耗时：%v ms", (p.EndUnixTime-p.StartUnixTime)/1000/1000)
	return p.Result
}

func (p *Base) Before() revel.Result {
	p.StartUnixTime = time.Now().UnixNano()

	switch p.Request.Method {
	case "GET":
		p.In = p.Params.Query
	case "POST":
		p.In = p.Params.Form
	}

	p.Cookie = map[string]string{}
	for _, item := range p.Request.Cookies() {
		p.Cookie[item.Name] = item.Value
	}

	remoteAddr := p.Request.RemoteAddr
	if value, ok := p.Request.Header["X-Real-Ip"]; ok {
		if len(value) > 0 {
			remoteAddr = value[0]
		}
	}

	global.Logger.Info(strings.Repeat("*", 100))
	global.Logger.Info("%v  %v  %v  %v", p.Request.Proto, p.Request.Method,
		p.Request.RequestURI, remoteAddr)
	global.Logger.Info("%v", p.Request.Header)
	global.Logger.Info("传入的Cookie:%v", p.Cookie)
	global.Logger.Info("传入参数:%v", p.In)

	return nil
}

func (p *Base) CheckParms(strParms map[string]string, intParms map[string]int, floatParms map[string]float64) error {
	err1 := p.CheckStrParms(strParms)
	err2 := p.CheckIntParms(intParms)
	err3 := p.CheckFloatParms(floatParms)
	if err1 == nil && err2 == nil && err3 == nil {
		return nil
	}
	msg := ""
	if err1 != nil {
		msg = msg + err1.Error() + ","
	}
	if err2 != nil {
		msg = msg + err2.Error() + ","
	}
	if err3 != nil {
		msg = msg + err3.Error() + ","
	}
	return errors.New(msg)
}

// 检查上传参数
func (p *Base) CheckStrParms(strParms map[string]string) error {
	if strParms == nil {
		return nil
	}
	lossKey := []string{}
	errmsg := ""
	for key, value := range strParms {
		if _, ok := p.In[key]; !ok {
			if value != define.PARMS_MUST_STR {
				delete(strParms, key)
			} else {
				lossKey = append(lossKey, key)
			}
		} else {
			strParms[key] = p.In[key][0]
		}
	}

	if len(lossKey) > 0 {
		errmsg = "缺少参数：" + strings.Join(lossKey, ",")
		return errors.New(errmsg)
	}

	return nil
}

func (p *Base) CheckIntParms(intParms map[string]int) error {
	if intParms == nil {
		return nil
	}

	lossKey := []string{}
	typeErrKey := []string{}
	errmsg := ""
	for key, value := range intParms {
		if _, ok := p.In[key]; !ok {
			if value != define.PARMS_MUST_INT {
				delete(intParms, key)
			} else {
				lossKey = append(lossKey, key)
			}
		} else {
			var err error
			intParms[key], err = strconv.Atoi(p.In[key][0])
			if err != nil {
				typeErrKey = append(typeErrKey, key)
			}
		}
	}

	if len(lossKey) > 0 || len(typeErrKey) > 0 {
		if len(lossKey) > 0 {
			errmsg = "缺少参数：" + strings.Join(lossKey, ",")
		}
		if len(typeErrKey) > 0 {
			errmsg = errmsg + "参数类型错误：" + strings.Join(typeErrKey, ",")
		}

		return errors.New(errmsg)
	}

	return nil
}

func (p *Base) CheckFloatParms(floatParms map[string]float64) error {
	if floatParms == nil {
		return nil
	}

	lossKey := []string{}
	typeErrKey := []string{}
	errmsg := ""
	for key, value := range floatParms {
		if _, ok := p.In[key]; !ok {
			if value != define.PARMS_MUST_FLOAT64 {
				delete(floatParms, key)
			} else {
				lossKey = append(lossKey, key)
			}
		} else {
			var err error
			floatParms[key], err = strconv.ParseFloat(p.In[key][0], 64)
			if err != nil {
				typeErrKey = append(typeErrKey, key)
			}
		}
	}

	if len(lossKey) > 0 || len(typeErrKey) > 0 {
		if len(lossKey) > 0 {
			errmsg = "缺少参数：" + strings.Join(lossKey, ",")
		}
		if len(typeErrKey) > 0 {
			errmsg = errmsg + "参数类型错误：" + strings.Join(typeErrKey, ",")
		}

		return errors.New(errmsg)
	}

	return nil
}

func (p *Base) Ret(ret int, msg string, result map[string]interface{}) revel.Result {
	p.Out.Ret = ret
	p.Out.Msg = msg
	p.Out.Data = result
	if result == nil {
		p.Out.Data = map[string]interface{}{}
	}
	return p.RenderJson(p.Out)
}

func (p *Base) SuccessRet(msg string, result map[string]interface{}) revel.Result {
	return p.Ret(define.RET_SUCCESS, msg, result)
}

func (p *Base) ErrParamsRet(msg string, result map[string]interface{}) revel.Result {
	return p.Ret(define.RET_PARMS_ERR, msg, result)
}

func (p *Base) ErrSystemRet(msg string, result map[string]interface{}) revel.Result {
	return p.Ret(define.RET_SYSTEM_ERR, msg, result)
}

type GrantBase struct {
	Base
}

func (p *GrantBase) isLogin() bool {
	return false
}

func (p *GrantBase) Before() revel.Result {
	if !p.isLogin() {
		return p.Ret(define.RET_RELOGIN, "身份过期，请重新登陆", nil)
	}
	return nil
}
