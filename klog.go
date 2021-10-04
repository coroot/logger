package logger

import (
	"fmt"
	"k8s.io/klog/v2"
)

func NewKlog(prefix string) *klogLogger {
	if prefix != "" {
		prefix += ": "
	}
	return &klogLogger{prefix: prefix}
}

type klogLogger struct {
	prefix string
}

func (l *klogLogger) Error(args ...interface{}) {
	klog.ErrorDepth(1, l.prefix, fmt.Sprintln(args...))
}

func (l *klogLogger) Errorf(format string, args ...interface{}) {
	klog.ErrorDepth(1, l.prefix, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Warning(args ...interface{}) {
	klog.WarningDepth(1, l.prefix, fmt.Sprintln(args...))
}

func (l *klogLogger) Warningf(format string, args ...interface{}) {
	klog.WarningDepth(1, l.prefix, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Info(args ...interface{}) {
	klog.InfoDepth(1, l.prefix, fmt.Sprintln(args...))
}

func (l *klogLogger) Infof(format string, args ...interface{}) {
	klog.InfoDepth(1, l.prefix, fmt.Sprintf(format, args...))
}
