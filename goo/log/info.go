package gooLog

import (
	"bytes"
	"fmt"
	"time"
)

type LogInfo struct {
	Level Level
	Data  []interface{}
	Trace []string
}

func (i *LogInfo) Bytes() []byte {
	var bf bytes.Buffer

	bf.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	bf.WriteString(fmt.Sprintf(" %s", LevelText[i.Level]))

	for _, msg := range i.Data {
		bf.WriteString(fmt.Sprintf(" %s", fmt.Sprint(msg)))
	}

	for _, msg := range i.Trace {
		bf.WriteString("\r\n")
		bf.WriteString(fmt.Sprintf(" %s", fmt.Sprint(msg)))
	}

	bf.WriteString("\r\n")

	return bf.Bytes()
}
