package commons

import (
	"os"
)

func IsOpenshiftHost() bool {
	return os.Getenv("OPENSHIFT_BUILD_NAME") != ""
}

func SearchEnv(envKey string, defaultValue string) (answer string) {
	value, exist := os.LookupEnv(envKey)
	if os.Getenv("OPENSHIFT_BUILD_NAME") != "" {
		if exist {
			return value
		} else {
			return "null"
		}
	} else {
		return defaultValue
	}
}
