package logger

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logFileName = "log"

var levelMap = map[string]zapcore.Level{
	"DEBUG":  zapcore.DebugLevel,
	"INFO":   zapcore.InfoLevel,
	"WARN":   zapcore.WarnLevel,
	"ERROR":  zapcore.ErrorLevel,
	"DPANIC": zapcore.DPanicLevel,
	"PANIC":  zapcore.PanicLevel,
	"FATAL":  zapcore.FatalLevel,
}

func checkConfig(config *Config) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}
	if len(config.LogPath) == 0 {
		return fmt.Errorf("logPath is nil")
	}
	if len(config.Level) == 0 {
		return fmt.Errorf("level is nil")
	}
	if config.MaxAge <= 0 {
		return fmt.Errorf("maxAge: %d must be positive", config.MaxAge)
	}
	if config.RotationSize <= 0 {
		return fmt.Errorf("rotationSize: %d must be positive", config.RotationSize)
	}
	if config.RotationTime <= 0 {
		return fmt.Errorf("rotationTime: %d must be positive", config.RotationTime)
	}
	return nil
}

func getWriter(config *Config) (zapcore.WriteSyncer, error) {
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
	return zapcore.AddSync(file), nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func converLevel(configLevel string) zapcore.Level {
	if zapLevel, ok := levelMap[configLevel]; ok {
		return zapLevel
	}
	return zapcore.InfoLevel
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
}
