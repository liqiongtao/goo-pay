package gooWeixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"googo.io/goo"
	gooHttp "googo.io/goo/http"
	gooUtils "googo.io/goo/utils"
)

// ---------------------------------
// -- 小程序登录
// ---------------------------------

type JsCode2SessionResponse struct {
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

func JsCode2Session(appid, secret, code string) (*JsCode2SessionResponse, error) {
	jscode2sess_url := fmt.Sprintf(sns_jsscode2sess_url, appid, secret, code)
	buf, err := gooHttp.NewRequest().Get(jscode2sess_url)
	if err != nil {
		return nil, err
	}

	rsp := &JsCode2SessionResponse{}
	if err := json.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != 0 {
		return nil, errors.New(rsp.Errmsg)
	}

	return rsp, nil
}

// ---------------------------------
// -- 解析用户数据
// ---------------------------------

type MinipUserInfoResponse struct {
	Openid   string `json:"openid"`
	Unionid  string `json:"unionid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatarUrl"`
	Gender   int    `json:"gender"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
}

func MinipUserInfo(sessionKey, encryptedData, iv string) (*MinipUserInfoResponse, error) {
	data := gooUtils.Base64Decode(encryptedData)
	key := gooUtils.Base64Decode(sessionKey)

	buf, err := gooUtils.Decrypt(data, key, gooUtils.Base64Decode(iv))
	if err != nil {
		return nil, err
	}

	userInfo := &MinipUserInfoResponse{}
	if err = json.Unmarshal(buf, userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

// ---------------------------------
// -- 发送模板消息
// ---------------------------------

func SendTemplateMessage(appid, secret, openid, templateId, page, formId string, data interface{}) error {
	accessToken, err := GetCGIAccessToken(appid, secret)
	if err != nil {
		return err
	}

	params := goo.Params{
		"access_token": accessToken,
		"touser":       openid,
		"template_id":  templateId,
		"page":         page,
		"form_id":      formId,
		"data":         data,
	}

	messageTplSendUrl := fmt.Sprintf(message_tpl_send_url, accessToken)
	buf, err := gooHttp.NewRequest().JsonContentType().Post(messageTplSendUrl, params.Json())
	if err != nil {
		return err
	}

	rst := &goo.Params{}
	if err := json.Unmarshal(buf, rst); err != nil {
		return err
	}
	if rst.GetInt("errcode") != 0 {
		return errors.New(rst.GetString("errmsg"))
	}

	return nil
}
