package tool

import (
	"sp-go-service/src/Resources/listing"
)

func GetErrorDetails(iderror string) (errcode int, reason string) {
	for _, errorInfo := range listing.ListingError {
		if errorInfo.Id == iderror {
			return errorInfo.Code, errorInfo.Reason
		}
	}
	return 999999, "Constant not found in ListingError"
}
