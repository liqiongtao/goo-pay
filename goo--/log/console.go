package gooLog

import (
	"io"
	"os"
)

type Console struct {
	Out io.Writer
}

func (c *Console) Init() {
	c.Out = os.Stdout
}

func (c *Console) Output(info *LogInfo) {
	c.Out.Write(info.Bytes())
}
