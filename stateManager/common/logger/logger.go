package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func SetLogLevel(level string) {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Printf("Invalid log level: %s. Using default level 'info'.\n", level)
		logrusLevel = logrus.InfoLevel
	}
	Log.SetLevel(logrusLevel)
}

func Init(level string) {
	SetLogLevel(level) // 초기 레벨을 "info"로 설정, 원하는 레벨로 변경 가능
	// 다른 로거 초기화 설정
}
