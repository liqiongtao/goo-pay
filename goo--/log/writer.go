package gooLog

type iWriter interface {
	Init()
	Output(info *LogInfo)
}
