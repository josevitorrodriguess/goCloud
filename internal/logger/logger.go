package logger

import (
	"fmt"
	"sync"
)

// LogLevel representa o nível do log
// Pode ser expandido futuramente

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var levelNames = [...]string{"DEBUG", "INFO", "WARN", "ERROR"}

// Logger estrutura principal
// Pode ser expandida para suportar saída em arquivo, formatação, etc.
type Logger struct {
	mu    sync.Mutex
	level LogLevel
}

var defaultLogger = &Logger{level: INFO}

// SetLevel permite alterar o nível global de log
func SetLevel(l LogLevel) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.level = l
}

// log é a função interna que faz o print
func (l *Logger) log(lv LogLevel, msg string, args ...interface{}) {
	if lv < l.level {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Printf("[%s] %s\n", levelNames[lv], fmt.Sprintf(msg, args...))
}

// Funções públicas para cada nível de log
func Info(msg string, args ...interface{}) {
	defaultLogger.log(INFO, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	defaultLogger.log(WARN, msg, args...)
}

func Error(msg string, args ...interface{}) {
	defaultLogger.log(ERROR, msg, args...)
}

func Debug(msg string, args ...interface{}) {
	defaultLogger.log(DEBUG, msg, args...)
}
