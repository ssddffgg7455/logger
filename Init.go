package logger

import (
	"fmt"
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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
	writer, err := GetWriter(config)
	if err != nil {
		return nil, fmt.Errorf("GetWriter err: %v", err)
	}
	writeSyncer := zapcore.AddSync(writer)
	encoder := getEncoder()
	level := converLevel(config.Level)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
	return sugarLogger, nil
}

// InitWithWriter 初始化，支持自定义输出 os.Stdout为标准输出
func InitWithWriter(configLevel string, writer zapcore.WriteSyncer) (*zap.SugaredLogger, error) {
	encoder := getEncoder()
	level := converLevel(configLevel)
	core := zapcore.NewCore(encoder, writer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
	return sugarLogger, nil
}

// GetWriter 根据配置获取writer 以供其他插件使用
func GetWriter(config *Config) (io.Writer, error) {
	err := checkConfig(config)
	if err != nil {
		return nil, err
	}
	file, err := rotatelogs.New(
		config.LogPath+"/"+logFileName+"-%Y%m%d%H.log",                            // 实际生成的文件名 log-YYmmddHH.log
		rotatelogs.WithMaxAge(time.Duration(config.MaxAge*24)*time.Hour),          // 最长保存MaxSaveDay天
		rotatelogs.WithRotationTime(time.Duration(config.RotationTime)*time.Hour), // 24小时切割一次
		rotatelogs.WithRotationSize(config.RotationSize*1024*1024),                // 分割日志的文件大小单位
	)
	if err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}
