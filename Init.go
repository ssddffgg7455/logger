package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config log 配置
type Config struct {
	Level        string // 日志分级 DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL
	LogPath      string // 日志路径
	MaxAge       int    // 最大保存时间 单位 day
	RotationSize int64  // 日志文件滚动size 单位 M
	RotationTime int    // 日志滚动周期 单位 hour
}

var sugarLogger *zap.SugaredLogger

// Init 初始化
func Init(config *Config) (*zap.SugaredLogger, error) {
	err := checkConfig(config)
	if err != nil {
		return nil, err
	}
	writeSyncer, err := getWriter(config)
	if err != nil {
		return nil, fmt.Errorf("getWriter err: %v", err)
	}
	encoder := getEncoder()
	level := converLevel(config.Level)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
	return sugarLogger, nil
}

// Init 初始化，支持自定义输出 os.Stdout为标准输出
func InitWithWriter(configLevel string, writer zapcore.WriteSyncer) (*zap.SugaredLogger, error) {
	encoder := getEncoder()
	level := converLevel(configLevel)
	core := zapcore.NewCore(encoder, writer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
	return sugarLogger, nil
}
