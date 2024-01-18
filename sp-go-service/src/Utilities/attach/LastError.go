package attach

import (
	"sp-go-service/src/Resources/global"
)

func LastError(errorcode int, reason string) {
	global.LastError["code"] = errorcode
	global.LastError["reason"] = reason
}
