package tool

import "sp-go-service/src/Resources/global"

func CleanGlobalEnvs() {
	global.RequestHeaders = make(map[string]string)
	global.ResponseBody = make(map[string]interface{})
	global.LastError = make(map[string]interface{})
}
