# 简单封装的日志工具

## 初始化

### 滚动日志配置
```
sugarLog, err := logger.Init(&logger.Config{
		Level:        level,        // 日志分级 DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL
		LogPath:      logPath,      // 日志路径
		MaxAge:       maxAge,       // 最大保存时间 单位 day
		RotationSize: rotationRize, // 日志文件滚动size 单位 M
		RotationTime: rotationTime, // 日志滚动周期 单位 hour
})
if err != nil {
    fmt.Printf("logger init err [%v] \n", err)
}
defer sugarLog.Sync()
```

### 自定义输出 os.Stdout为标准输出
```
sugarLog, err := logger.InitWithWriter(level, os.Stdout)
if err != nil {
    fmt.Printf("logger init err [%v] \n", err)
}
defer sugarLog.Sync()
```

## 日志打印
```
ctx := logger.CtxWithTraceId() // 生成携带trace_id的上下文
logger.Debugw(ctx, msg, keysAndValues ...)
logger.Infow(ctx, msg, keysAndValues ...)
logger.Warnw(ctx, msg, keysAndValues ...)
logger.Errorw(ctx, msg, keysAndValues ...)
logger.Fatalw(ctx, msg, keysAndValues ...)
logger.ErrorwWithTrace(ctx, msg, keysAndValues ...)
```
