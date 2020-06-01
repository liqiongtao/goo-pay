package gooWeixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"googo.io/goo"
	gooHttp "googo.io/goo/http"
	gooLog "googo.io/goo/log"
	"net/url"
	"strings"
	"time"
)

// 获取H5授权链接
// authorizeUrl: 授权地址
// originUrl: 授权成功后的回跳地址
func Oauth2AuthorizeUrl(appid, authorizeUrl, originUrl string) string {
	var redirectUrl string
	if strings.Index(authorizeUrl, "?") != -1 {
		redirectUrl = url.QueryEscape(authorizeUrl + "&redirect_url=" + url.QueryEscape(originUrl))
	} else {
		redirectUrl = url.QueryEscape(authorizeUrl + "?redirect_url=" + url.QueryEscape(originUrl))
	}
	oauth2Url := fmt.Sprintf(oauth2_authorize_url, appid, redirectUrl, state)
	gooLog.Debug("wx-h5-oauth2", goo.Params{
		"authorizeUrl": authorizeUrl,
		"originUrl":    originUrl,
		"redirectUrl":  redirectUrl,
		"oauth2Url":    oauth2Url,
	})
	return oauth2Url
}

type Oauth2AccessTokenResponse struct {
	AccessToken  string        `json:"access_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
	RefreshToken string        `json:"refresh_token"`
	Openid       string        `json:"openid"`
	Unionid      string        `json:"unionid"`
	Scope        string        `json:"scope"`
	Errcode      int           `json:"errcode"`
	Errmsg       string        `json:"errmsg"`
}

func Oauth2AccessToken(appid, secret, code string) (*Oauth2AccessTokenResponse, error) {
	accessTokenUrl := fmt.Sprintf(sns_oauth2_accessToken_url, appid, secret, code)
	buf, err := gooHttp.NewRequest().Get(accessTokenUrl)
	if err != nil {
		return nil, err
	}

	rsp := &Oauth2AccessTokenResponse{}
	if err := json.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	if rsp.Errcode != 0 {
		return nil, errors.New(rsp.Errmsg)
	}

	return rsp, nil
}

type SnsUserInfoResponse struct {
	Openid   string `json:"openid"`
	Unionid  string `json:"unionid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"headimgurl"`
	Gender   int    `json:"sex"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
}

func SnsUserInfo(accessToken, openid string) (*SnsUserInfoResponse, error) {
	userInfoUrl := fmt.Sprintf(sns_userinfo_url, accessToken, openid)
	buf, err := gooHttp.NewRequest().Get(userInfoUrl)
	if err != nil {
		return nil, err
	}

	userInfo := &SnsUserInfoResponse{}
	if err := json.Unmarshal(buf, userInfo); err != nil {
		return nil, err
	}
	if userInfo.Errcode != 0 {
		return nil, errors.New(userInfo.Errmsg)
	}

	return userInfo, nil
}
