package logger

import (
	"fmt"
	"log"
	"os"
	"sp-go-service/src/Resources/global"
	"sp-go-service/src/Resources/parameter"
	"sp-go-service/src/Utilities/attach"
	"sp-go-service/src/Utilities/extract"
	"sp-go-service/src/Utilities/tool"
	"strconv"
	"strings"
	"time"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", 0)
}

func Console(level string, iderror string, exception string) {

	inLevel := strings.ToUpper(level)

	if iderror == "" {
		iderror = "ERROR_UNKNOWN"
	}

	var errcode int
	var mreason string
	var infodebug bool
	if inLevel == "INFO" || inLevel == "DEBUG" {
		errcode, mreason = 0, "Notification:"
		iderror = "ONLY_INFO"
		infodebug = true
	} else {
		errcode, mreason = tool.GetErrorDetails(iderror)
	}

	paramLevel := strings.ToUpper(parameter.LOG_LEVEL)
	var runLog bool
	switch {
	case (paramLevel == "INFO") && (inLevel == "INFO"):
		runLog = true

	case (paramLevel == "WARNING") && (inLevel == "WARNING"):
		runLog = true

	case (paramLevel == "ERROR") && (inLevel == "ERROR"):
		runLog = true

	case (paramLevel == "DEBUG") && (inLevel == "INFO" || inLevel == "WARNING" || inLevel == "ERROR" || inLevel == "DEBUG"):
		runLog = true

	default:
		runLog = false
	}

	timestamp := time.Now().Format(time.RFC3339)
	funcCall, filePath, fileLine, pid := tool.GetFunction(2)

	var registry string
	if regis := extract.RequestHeaders(); regis != "" {
		registry = fmt.Sprintf("-> [%v]", extract.RequestHeaders())
	} else {
		registry = ""
	}

	reason := fmt.Sprintf("-> [%v %v] [Function:%v FilePath:%v FileLine:%v PID:%v]", mreason, exception, funcCall, filePath, fileLine, pid)
	detailfunc := fmt.Sprintf("-> [Function:%v FilePath:%v FileLine:%v PID:%v]", funcCall, filePath, fileLine, pid)

	if runLog {
		logger.Printf("%s %s %s --- [%s] [%s:%s] %s %s %s \n\n", timestamp, inLevel, strconv.Itoa(pid), funcCall, filePath, strconv.Itoa(fileLine), registry, reason, detailfunc)
	}

	attach.History(errcode, reason)

	if !infodebug {
		attach.LastError(errcode, reason)
	}

	if inLevel == "ERROR" {
		global.ResponseBody = make(map[string]interface{})
	}

}
