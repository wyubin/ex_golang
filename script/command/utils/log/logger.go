// add logger with zap
package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(logLevel zapcore.Level, ioLogs ...io.Writer) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	var cores []zapcore.Core = []zapcore.Core{
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
	}
	for _, ioLog := range ioLogs {
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(ioLog), logLevel))
	}
	core := zapcore.NewTee(cores...)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
