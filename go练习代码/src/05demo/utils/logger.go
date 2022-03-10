package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() () {
	Encoder := getEncoder()
	writeSyncer := getWriteSyncer()
	core := zapcore.NewCore(Encoder, writeSyncer, zapcore.DebugLevel)
	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
}

func getEncoder() zapcore.Encoder {
	newEncoder := zap.NewProductionEncoderConfig()
	newEncoder.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(newEncoder)
}

func getWriteSyncer() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.max_size"),
		MaxAge:     viper.GetInt("log.max_day"),
		MaxBackups: viper.GetInt("log.max_backups"),
	}
	return zapcore.AddSync(file)
}
