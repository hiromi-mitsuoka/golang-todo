package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// 読み書き、作成、追記
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// TODO: 内容を理解する
	// 書き込み方法、書き込み先指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// Specify format of log
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// Specify log output destination
	log.SetOutput(multiLogFile)
}
