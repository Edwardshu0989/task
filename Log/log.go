package Log

import (
	"github.com/sirupsen/logrus"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"os"
	"time"
)

var Log *logrus.Logger

func init(){
	//创建log目录
	os.MkdirAll("./logs",0777)
	Log =logrus.New()
	Log.SetLevel(logrus.DebugLevel)
	infoWriter,_:=rotatelogs.New(
		"./logs/info.%Y%m%d.log",
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		)
	writeMap:=lfshook.WriterMap{
		logrus.DebugLevel:infoWriter,
		logrus.ErrorLevel:infoWriter,
		logrus.InfoLevel:infoWriter,
		logrus.FatalLevel:infoWriter,
	}

	Log.Hooks.Add(lfshook.NewHook(writeMap,&logrus.TextFormatter{
		TimestampFormat:           "2020-07-30 02:03:04",
		DisableSorting:            false,
	}))
}

func Info(arg ...interface{}){
	Log.Info(arg...)
}
