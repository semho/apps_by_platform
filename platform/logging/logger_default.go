package logging

import (
	"fmt"
	"log"
)

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLogger) Trace(msc string) {
	l.write(Trace, msc)
}

func (l *DefaultLogger) Tracef(template string, vals ...any) {
	l.write(Trace, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Debug(msc string) {
	l.write(Debug, msc)
}

func (l *DefaultLogger) Debugf(template string, vals ...any) {
	l.write(Debug, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Info(msc string) {
	l.write(Information, msc)
}

func (l *DefaultLogger) Infof(template string, vals ...any) {
	l.write(Information, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Warn(msc string) {
	l.write(Warning, msc)
}

func (l *DefaultLogger) Warnf(template string, vals ...any) {
	l.write(Warning, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Panic(msc string) {
	l.write(Fatal, msc)
	if l.triggerPanic {
		panic(msc)
	}
}

func (l *DefaultLogger) Panicf(template string, vals ...any) {
	formattedMsg := fmt.Sprintf(template, vals...)
	l.write(Fatal, formattedMsg)
	if l.triggerPanic {
		panic(formattedMsg)
	}
}
