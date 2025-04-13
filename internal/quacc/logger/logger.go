package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	info  *log.Logger
	debug *log.Logger
)

func Info(msg ...string) {
	if info != nil {
		info.Println(msg)
	}
}

func init() {
	file, err := os.OpenFile("quacc.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("failed to create log file")
	}

	info = log.New(file, "[INFO] \t", log.LstdFlags)
}
