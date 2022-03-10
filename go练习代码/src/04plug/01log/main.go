package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	//"os"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

var logger *zap.Logger
var sugerlogger *zap.SugaredLogger

//打印到终端
//func InitLog() {
//	logger, _ = zap.NewProduction()
//	sugerlogger = logger.Sugar()
//}

//打印到文件
func InitLog() {
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
	sugerlogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(encoder)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./test.log")
	file := &lumberjack.Logger{
		Filename:   "./test.log", //文件名
		MaxSize:    10,           //文件大小
		MaxBackups: 5,            //备份数量
		MaxAge:     30,           //最大备份天数
		Compress:   false,        //是否压缩
	}
	return zapcore.AddSync(file)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugerlogger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		sugerlogger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

//func main() {
//	InitLog()
//	defer logger.Sync()
//	simpleHttpGet("www.google.com")
//	simpleHttpGet("http://www.baidu.com")
//}
func main() {
	InitLog()
	defer logger.Sync()
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/hello", func(c *gin.Context) {
		logger.Info("___________________________________")
		c.JSON(http.StatusOK, gin.H{"data": "hello"})
	})
	r.Run()
}
