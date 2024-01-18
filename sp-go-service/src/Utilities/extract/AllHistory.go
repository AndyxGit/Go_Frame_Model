package extract

import (
	"sp-go-service/src/Resources/global"
)

func AllHistory() map[string][]map[string]interface{} {

	allHistoryCopy := make(map[string][]map[string]interface{})

	for sessionID, details := range global.HistoryMap {
		var sessionHistory []map[string]interface{}
		for _, detail := range details {
			detailMap := map[string]interface{}{
				"code":   detail.Code,
				"reason": detail.Reason,
			}
			sessionHistory = append(sessionHistory, detailMap)
		}
		allHistoryCopy[sessionID] = sessionHistory
	}

	return allHistoryCopy
}
