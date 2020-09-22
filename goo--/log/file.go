package gooLog

import (
	"fmt"
	"io"
	"os"
	"time"
)

type File struct {
	Dir         string
	FileName    string
	dateLogFile string
	out         io.Writer
}

func (f *File) Init() {
	if f.Dir == "" {
		f.Dir = "logs/"
	}

	if _, err := os.Stat(f.Dir); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(f.Dir, 0755)
		}
	}
}

func (f *File) Output(info *LogInfo) {
	if dateLog := f.dateLog(); f.dateLogFile != dateLog {
		f.dateLogFile = dateLog
		f.out, _ = os.OpenFile(f.Dir+f.dateLogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}

	f.out.Write(info.Bytes())
}

func (f *File) dateLog() string {
	if f.FileName == "" {
		return fmt.Sprintf("%s.log", time.Now().Format("20060102"))
	}
	return fmt.Sprintf("%s_%s.log", f.FileName, time.Now().Format("20060102"))
}
