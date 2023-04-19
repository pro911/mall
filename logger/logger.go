package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"mall/settings"
	"os"
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int, closeStdout bool) zapcore.WriteSyncer {
	// 利用io.MultiWriter支持文件和终端两个输出目标
	var ws io.Writer

	//日志切割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		LocalTime:  false,
		Compress:   false,
	}
	ws = lumberjackLogger
	if closeStdout {
		ws = io.MultiWriter(lumberjackLogger, os.Stdout)
	}
	return zapcore.AddSync(ws)
}

func Init(cfg *settings.LogConfig) (err error) {
	encoder := getEncoder()

	writerSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
		cfg.CloseStdout,
	)

	//创建一个错误的日志收集
	writerSyncerErr := getLogWriter(
		cfg.ErrFilename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
		cfg.CloseStdout,
	)

	var l = new(zapcore.Level)
	if err = l.UnmarshalText([]byte(cfg.Level)); err != nil {
		return err
	}

	core := zapcore.NewCore(encoder, writerSyncer, l)
	coreErr := zapcore.NewCore(encoder, writerSyncerErr, zapcore.ErrorLevel)

	//使用NewTee将多个文件合并到core
	lg := zap.New(zapcore.NewTee(core, coreErr), zap.AddCaller())
	//替换zap库中全局的logger
	zap.ReplaceGlobals(lg)
	return
}
