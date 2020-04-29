package logger

/*
该log包依赖gin.Context, 便于在一个request生命周期内的log都带上唯一请求ID，故需要初始化一个生成全局中间件以生成唯一请求ID
不需要记录请求ID的时候ctx传nil即可
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"ppap/backup/go/config"
)

var defaultLogger *zap.SugaredLogger

// Setup 初始化logger
func SetUp() {
	defaultLogger = initLogger().Sugar()
	defaultLogger.Info("init logger success")
}

func initLogger() *zap.Logger {
	var level zapcore.LevelEnabler
	switch config.Get("log.level") {
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.DebugLevel
	}

	return zap.New(zapcore.NewCore(
		// json格式
		getEncoder(),
		// log文件输出
		GetLogWriter(),
		// 日志级别
		level,
	))
}

// 配置日志输出格式
func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// json 格式
	return zapcore.NewJSONEncoder(encodeConfig)
}

// 配置日志文件
func GetLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Get("log.filename"),
		MaxSize:   	config.GetInt("log.max_size"),
		MaxBackups: config.GetInt("log.max_backups"),
		MaxAge:     config.GetInt("log.max_age"),
		Compress:   false,
	})
}

// getLogger 每一个请求生成一个带请求ID的logger
func getLogger(ctx *gin.Context) *zap.SugaredLogger {
	if ctx == nil {
		return defaultLogger
	}
	return defaultLogger.With(zap.Int64(XRI, getRequestID(ctx)))
}

// Info info
func Info(ctx *gin.Context, msg string, keyAndValues ...interface{}) {
	getLogger(ctx).Infow(msg, keyAndValues...)
}

// Debug debug
func Debug(ctx *gin.Context, msg string, keyAndValues ...interface{}) {
	getLogger(ctx).Debugw(msg, keyAndValues...)
}

// Warn warn， 顺便发送sentry
func Warn(ctx *gin.Context, msg string, keyAndValues ...interface{}) {
	getLogger(ctx).Warnw(msg, keyAndValues...)
}

// Error error, 顺便发送sentry
func Error(ctx *gin.Context, msg string, keyAndValues ...interface{}) {
	getLogger(ctx).Errorw(msg, keyAndValues...)
}
