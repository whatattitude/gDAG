// Convenient log library out of the box
package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/whatattitude/gDAG/lib/env"
	"github.com/whatattitude/gDAG/lib/file"
)

var Logger = DefaultInitLogger("")

func DefaultInitLogger(path string) *zap.Logger {
	if path == "" {
		path = "./log/statistics/"
	}

	goEnv, err := env.GetEnvDefault()
	if err != nil {
		fmt.Println("get go env error. using default logger path")
	}
	path = strings.Replace(goEnv.Gomod, "go.mod", "log/statistics/", 1)
	return InitLogger(path)
}

func InitLogger(logFilePath string) *zap.Logger {

	file.CreateDir(logFilePath)
	// 设置日志输出格式为JSON (参数复用NewDevelopmentEncoderConfig)
	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	infoFile := logFilePath + "info.log"
	errFile := logFilePath + "error.log"
	warnFile := logFilePath + "warnning.log"
	debugFile := logFilePath + "debug.log"

	zcore := zapcore.NewTee(
		zapcore.NewCore(encoder, getLumberjackConfig(infoFile), zap.DebugLevel),
		zapcore.NewCore(encoder, getLumberjackConfig(errFile), zap.ErrorLevel),
		zapcore.NewCore(encoder, getLumberjackConfig(warnFile), zap.WarnLevel),
		zapcore.NewCore(encoder, getLumberjackConfig(debugFile), zap.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	// 创建日志记录器
	logger := zap.New(zcore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	defer logger.Sync()
	return logger

}

func getLumberjackConfig(fileName string) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName, //日志文件
		MaxSize:    10,       //单文件最大容量(单位MB)
		MaxBackups: 3,        //保留旧文件的最大数量
		MaxAge:     1,        // 旧文件最多保存几天
		Compress:   false,    // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberjackLogger)
}
