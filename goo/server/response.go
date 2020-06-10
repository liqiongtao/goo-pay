package gooServer

import (
	"encoding/json"
)

type Response struct {
	Status  int         `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (rsp Response) ToString() string {
	buf, err := json.Marshal(&rsp)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}
