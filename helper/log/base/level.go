package base

//日志输出级别
type LogLevel uint8

const (
	LEVEL_DEBUG LogLevel = iota + 1
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
	LEVEL_PANIC
)
