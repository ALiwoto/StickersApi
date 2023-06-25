package logging

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SUGARED *zap.SugaredLogger

func InitZapLog(debug bool) *zap.Logger {
	var config zap.Config
	if debug {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, _ := config.Build(zap.AddCallerSkip(1))
	return logger
}

func LoadLogger(debug bool) func() {
	if SUGARED != nil {
		return nil
	}
	loggerMgr := InitZapLog(debug)
	zap.ReplaceGlobals(loggerMgr)
	SUGARED = loggerMgr.Sugar()

	return func() {
		_ = loggerMgr.Sync()
	}
}

func Warn(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Warn(args...)
	} else {
		log.Println(args...)
	}
}

func Error(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Error(args...)
	} else {
		log.Println(args...)
	}
}

func Info(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Info(args...)
	} else {
		log.Println(args...)
	}
}

func Infof(template string, args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Infof(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Debug(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Debug(args...)
	} else {
		log.Println(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Debugf(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Fatal(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Fatal(args...)
	} else {
		log.Fatal(args...)
	}
}
