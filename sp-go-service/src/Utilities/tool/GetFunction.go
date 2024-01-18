package tool

import (
	"os"
	"runtime"
)

func GetFunction(position int) (functionName string, functionPath string, functionFileLine int, pid int) {
	pc, _, _, _ := runtime.Caller(position)
	funcCall := runtime.FuncForPC(pc).Name()
	filePath, fileLine := runtime.FuncForPC(pc).FileLine(pc)
	vpid := os.Getpid()
	return funcCall, filePath, fileLine, vpid
}
