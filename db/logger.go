package db

import (
	log "github.com/sirupsen/logrus"
	"xorm.io/core"
)

type Logger struct {
	*log.Entry
	level core.LogLevel
}

func (l *Logger) SetLevel(level core.LogLevel) {
	l.level = level
}

func (l *Logger) Level() core.LogLevel {
	return l.level
}

func (l *Logger) ShowSQL(show ...bool) {}
func (l *Logger) IsShowSQL() bool      { return false }
