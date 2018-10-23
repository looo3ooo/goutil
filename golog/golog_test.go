package log

import (
	"gogs.163.com/feiyu/goutil/golog/conf"
	"testing"
	"time"
)

func Test_logge(t *testing.T) {
	SetLogger(ZAPLOG,
		conf.WithLogType(conf.LogNormalType),
		conf.WithProjectName("a"))

	SetLogLevel(conf.ErrorLevel)
	Debug("this is zap")
	Debug("this is zap")
	SetLogLevel(conf.DebugLevel)
	Debug("this is zap")
	Debug("this is zap")

	time.Sleep(time.Second * 5)

}
