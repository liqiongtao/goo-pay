package gooSms

type ISms interface {
	Init()
	Send(mobile, action string) (string, error)
	Verify(mobile, action, code string) error
}

var __sms ISms

func Init(sms ISms) {
	__sms = sms
	__sms.Init()
}
