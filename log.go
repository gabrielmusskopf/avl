package avl

import "fmt"

const (
	NONE  = "none"
	DEBUG = "debug"
	INFO  = "info"
)

var LogLevel string

func IsDebug() bool {
    return LogLevel == DEBUG
}

func ToggleDebug() {
    if LogLevel == DEBUG {
        LogLevel = NONE
    } else {
        LogLevel = DEBUG
    }
}

func Info(s string, v ...any) {
	if LogLevel == INFO {
		str := fmt.Sprintf("[info] %s", s)
		fmt.Printf(str, v...)
	}
}

func Debug(s string, v ...any) {
	if LogLevel == DEBUG {
		str := fmt.Sprintf("[debug] %s", s)
		fmt.Printf(str, v...)
	}
}
