//+build debug

package logging

import (
	"log"
	"sync"
	"strings"
)

func Info(format string, v ...interface{}) {
	log.Printf(format, v...)
}

var GLOBAL_COMMAND_ID_COUNTER uint64 = 0
var GLOBAL_IPC_MUTEX sync.Mutex = sync.Mutex{}

func IPCCmdID() uint64 {
	GLOBAL_IPC_MUTEX.Lock()
	GLOBAL_COMMAND_ID_COUNTER++
	i := GLOBAL_COMMAND_ID_COUNTER
	GLOBAL_IPC_MUTEX.Unlock()
	return i
}

func ProcessStr(str string, maxLen int) string {
	if len(str) > maxLen {
		str = str[:maxLen] + "..."
	}
	return strings.ReplaceAll(str, "\n", "//")
}