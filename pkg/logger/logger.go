package logger

import (
	"food-recipes-backend/pkg/setting"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	*zap.Logger
}

func NewLogger(logConfig setting.LoggerSetting) *ZapLogger {
	logLevel := logConfig.LogLevel
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	case "panic":
		level = zap.PanicLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   logConfig.FileName,
		MaxSize:    logConfig.MaxAge, // megabytes
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge, //days
		Compress:   logConfig.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)
	return &ZapLogger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 1725891083.7075524 -> 2024-09-09T21:11:23.706+0700 
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encoderConfig.TimeKey = "time"
	// from "info" -> "INFO"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// caller -> file:line
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}