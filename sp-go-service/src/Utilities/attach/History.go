package attach

import (
	"sp-go-service/src/Resources/global"
	"sp-go-service/src/Resources/structure"
)

func History(errorcode int, reason string) {
	sessionID := global.RequestHeaders["SessionID"]
	global.HistoryMap[sessionID] = append(global.HistoryMap[sessionID], structure.ServiceDetail{Code: errorcode, Reason: reason})
}
