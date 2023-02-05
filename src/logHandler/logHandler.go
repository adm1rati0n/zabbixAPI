package logHandler

import (
	"log"
	"os"
	"time"
)

var File *os.File

func CreateFile() {
	dir, _ := os.Getwd()
	if _, err := os.Stat(dir + "/logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
	}

	file, err := os.Create(dir + "/logs/" + time.Now().Format("2006.01.02(15h04m05s)") + ".txt")
	File = file
	if err != nil {
		log.Fatal(err)
	}
}

func WriteLogs(text string) {
	_, err := File.WriteString("\n" + time.Now().Format("2006.01.02(15h04m05s)") + ": " + text)
	if err != nil {
		log.Fatal(err)
	}
	//defer File.Close()
}
