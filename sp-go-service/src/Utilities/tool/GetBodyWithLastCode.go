package tool

import (
	"sp-go-service/src/Resources/global"
	"sp-go-service/src/Utilities/extract"
)

func GetBodyWithLastCode() {

	lastDetail := extract.LastError()

	if len(lastDetail) == 0 {
		lastDetail["code"] = "0"
		lastDetail["reason"] = "OK"
	}

	if lastDetail["level"] == "ERROR" {
		global.ResponseBody = make(map[string]interface{})
	}

	global.ResponseBody["code"] = lastDetail["code"]
	global.ResponseBody["reason"] = lastDetail["reason"]

}
