package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
	if len(config.LogFile) == 0 {
		return fmt.Errorf("LogFile is nil")
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
