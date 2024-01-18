package extract

import (
	"sp-go-service/src/Resources/global"
)

func History() []map[string]interface{} {
	sessionID := global.RequestHeaders["SessionID"]

	var vcopy []map[string]interface{}

	for _, detail := range global.HistoryMap[sessionID] {
		detailMap := map[string]interface{}{
			"code":   detail.Code,
			"reason": detail.Reason,
		}
		vcopy = append(vcopy, detailMap)
	}

	return vcopy
}
