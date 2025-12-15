package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger
var SugarLog *zap.SugaredLogger

func Init() {
	writeSyncer := getWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	Log = zap.New(core, zap.AddCaller())
	// SugarLog = Log.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}
func getWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./blog.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)
	return zapcore.AddSync(file)
}
