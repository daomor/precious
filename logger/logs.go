package logger

import (
	"os"
	"log"
)

type MyLog struct {
	log *os.File
	Name string
}

func (l *MyLog) Message(message string) {
	log.SetOutput(l.log)
	log.Println(message)
}