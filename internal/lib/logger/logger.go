package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	*log.Logger
}

func New() *Logger {
	return &Logger{
		log.Default(),
	}
}

func (lgr *Logger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	lgr.Printf("[Error]: %s\n", msg)
}

func (lgr *Logger) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	lgr.Printf("[Info]: %s\n", msg)
}
