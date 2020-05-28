package gooConfig

import (
	gooLog "googo.io/goo/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadFile(filename string, config interface{}) error {
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		gooLog.Error(err.Error())
		return err
	}

	err = yaml.Unmarshal(bts, config)
	if err != nil {
		gooLog.Error(err.Error())
	}
	return err
}
