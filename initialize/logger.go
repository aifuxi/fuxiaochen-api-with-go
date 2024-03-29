package initialize

import (
	"fuxiaochen-api-with-go/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func SetupLogger() {
	writeSyncer := newLogWriter()
	encoder := newEncoder()
	consoleEncoder := newConsoleEncoder()
	c1 := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	c2 := zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zapcore.DebugLevel)

	var logger *zap.Logger

	if global.Conf.AppConfig.Mode == gin.DebugMode {
		core := zapcore.NewTee(c1, c2)
		logger = zap.New(core, zap.AddCaller())
	} else {
		logger = zap.New(c1, zap.AddCaller())
	}

	global.Logger = logger
	global.L = logger.Sugar()
}

func newEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func newConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func newLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   global.Conf.LogConfig.Filepath,
		MaxSize:    global.Conf.LogConfig.MaxSize,
		MaxBackups: global.Conf.LogConfig.MaxBackups,
		MaxAge:     global.Conf.LogConfig.MaxAge,
		Compress:   global.Conf.LogConfig.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
