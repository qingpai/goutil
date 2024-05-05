package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var lumberJackLogger lumberjack.Logger

func Debug(msg string) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelDebug, msg, pcs[0])
	debugLogger.Handler().Handle(context.Background(), r)
	fileLogger.Handler().Handle(context.Background(), r)
}
func Debugf(msg string, args ...any) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelDebug, fmt.Sprintf(msg, args...), pcs[0])
	debugLogger.Handler().Handle(context.Background(), r)
	fileLogger.Handler().Handle(context.Background(), r)
}

func Info(msg string) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, pcs[0])
	debugLogger.Handler().Handle(context.Background(), r)
	fileLogger.Handler().Handle(context.Background(), r)
}
func Infof(msg string, args ...any) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(msg, args...), pcs[0])
	debugLogger.Handler().Handle(context.Background(), r)
	fileLogger.Handler().Handle(context.Background(), r)
}

func Warn(msg string) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelWarn, msg, pcs[0])
	debugLogger.Handler().Handle(context.Background(), r)
	fileLogger.Handler().Handle(context.Background(), r)
}
func Warnf(msg string, args ...any) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelWarn, fmt.Sprintf(msg, args...), pcs[0])
	debugLogger.Handler().Handle(context.Background(), r)
	fileLogger.Handler().Handle(context.Background(), r)
}

func Error(msg string) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0])
	debugLogger.Handler().Handle(context.Background(), r)
	fileLogger.Handler().Handle(context.Background(), r)
}
func Errorf(msg string, args ...any) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(msg, args...), pcs[0])
	debugLogger.Handler().Handle(context.Background(), r)
	fileLogger.Handler().Handle(context.Background(), r)
}

var debugLogger *slog.Logger
var fileLogger *slog.Logger

func Init(level string) {
	var logLevel slog.LevelVar
	switch level {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "warn":
		logLevel.Set(slog.LevelWarn)
	default:
		logLevel.Set(slog.LevelError)
	}

	debugLogger = _initDebugLogger(&logLevel)
	fileLogger = _initFileLogger(&logLevel)
}

func _initDebugLogger(logLevel *slog.LevelVar) *slog.Logger {
	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       logLevel,
		ReplaceAttr: nil,
	}

	return slog.New(slog.NewTextHandler(os.Stderr, &opts))
}

func _initFileLogger(logLevel *slog.LevelVar) *slog.Logger {
	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       logLevel,
		ReplaceAttr: nil,
	}

	lumberJackLogger = lumberjack.Logger{
		Filename:   "logs/log.log",
		MaxSize:    5,
		MaxBackups: 0,
		MaxAge:     0,
		Compress:   true,
	}

	return slog.New(slog.NewJSONHandler(&lumberJackLogger, &opts))
}

func Rotate() {
	_ = lumberJackLogger.Rotate()
}
