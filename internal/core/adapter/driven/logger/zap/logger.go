package zap

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(logger *zap.Logger) Logger {
	return Logger{
		logger: logger.WithOptions(zap.AddCallerSkip(1)).Sugar(),
	}
}

func (l Logger) With(args ...interface{}) service.Logger {
	return Logger{
		logger: l.logger.With(args...),
	}
}

func (l Logger) Debug(msg string, args ...interface{}) {
	l.logger.Debugw(msg, args...)
}

func (l Logger) Info(msg string, args ...interface{}) {
	l.logger.Infow(msg, args...)
}

func (l Logger) Warn(msg string, args ...interface{}) {
	l.logger.Warnw(msg, args...)
}

func (l Logger) Error(msg string, args ...interface{}) {
	l.logger.Errorw(msg, args...)
}

func (l Logger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatalw(msg, args...)
}
