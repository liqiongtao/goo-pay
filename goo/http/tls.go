package gooHttp

import (
	"crypto/tls"
	gooLog "goo/log"
	"io/ioutil"
)

type Tls struct {
	CaCrtFile     string
	ClientCrtFile string
	ClientKeyFile string
}

func (this *Tls) CaCrt() []byte {
	if this.CaCrtFile == "" {
		return getCaCert()
	}
	bts, err := ioutil.ReadFile(this.CaCrtFile)
	if err != nil {
		gooLog.Error(err.Error())
	}
	return bts
}

func (this *Tls) ClientCrt() tls.Certificate {
	crt, err := tls.LoadX509KeyPair(this.ClientCrtFile, this.ClientKeyFile)
	if err != nil {
		gooLog.Error(err.Error())
	}
	return crt
}
