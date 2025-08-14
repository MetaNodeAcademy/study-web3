package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"task4/internal/config"
)

var Logger *zap.Logger

func InitLogger() error {
	// 配置日志输出位置和滚动策略
	hook := lumberjack.Logger{
		Filename:   config.GetConfig().Log.Path, // 日志文件路径
		MaxSize:    128,                         // 单个文件最大容量(MB)
		MaxBackups: 30,                          // 最多保留备份数
		MaxAge:     7,                           // 最多保留天数
		Compress:   true,                        // 是否压缩
	}

	// 设置日志级别
	var levelEnab zapcore.Level
	switch config.GetConfig().Log.Level {
	case "debug":
		levelEnab = zapcore.DebugLevel
	case "info":
		levelEnab = zapcore.InfoLevel
	case "warn":
		levelEnab = zapcore.WarnLevel
	case "error":
		levelEnab = zapcore.ErrorLevel
	default:
		levelEnab = zapcore.InfoLevel
	}

	// 配置编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 级别大写

	// 配置核心
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(&hook),
			levelEnab,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			levelEnab,
		),
	)

	// 创建日志实例
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return nil
}
