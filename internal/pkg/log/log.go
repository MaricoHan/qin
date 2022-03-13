package log

import (
	"io"
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file := "./log.log"
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0766)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(os.Stderr, logFile)

	Debug = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(multiWriter, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}
