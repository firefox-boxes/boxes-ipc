//+build !debug

package logging

func Info(format string, v ...interface{}) {}

func IPCCmdID() uint64 { return 1 }

func ProcessStr(str string, maxLen int) string { return str }