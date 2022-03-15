package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	multiLogfile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogfile)

}

// log.Lshortfile

// こやつは、log.Printlnをどこで発動させたのかを記してくれる

// multiLogfile := io.MultiWriter(os.Stdout, logfile)
// こやつはlogの出力先を複数設定したい時に使う
// 今回は、標準出力とlogfileのどっちにも出力する
