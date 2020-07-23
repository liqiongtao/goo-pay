package gooSms

import (
	"errors"
	"regexp"
)

func Send(mobile, action string) (string, error) {
	if mobile == "" {
		return "", errors.New("mobile is null")
	}
	if regexp.MustCompile(`^1[3,4,5,7,8]\d{9}$`).MatchString(mobile) == false {
		return "", errors.New("invalid mobile")
	}
	if action == "" {
		return "", errors.New("action is null")
	}
	return __sms.Send(mobile, action)
}

func Verify(mobile, action, code string) error {
	if mobile == "" {
		return errors.New("mobile is null")
	}
	if regexp.MustCompile(`^1[3,4,5,7,8]\d{9}$`).MatchString(mobile) == false {
		return errors.New("invalid mobile")
	}
	if action == "" {
		return errors.New("action is null")
	}
	if code == "" {
		return errors.New("code is null")
	}
	return __sms.Verify(mobile, action, code)
}
