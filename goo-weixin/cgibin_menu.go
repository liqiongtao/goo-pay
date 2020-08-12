package gooWeixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"googo.io/goo"
	gooHttp "googo.io/goo/http"
)

func MenuCreate(appid, secret, content string) error {
	accessToken := CGIToken(appid, secret).Get()

	menuCreateUrl := fmt.Sprintf(menu_create_url, accessToken)
	buf, err := gooHttp.NewRequest().JsonContentType().Post(menuCreateUrl, []byte(content))
	if err != nil {
		return err
	}

	rst := goo.Params{}
	if err := json.Unmarshal(buf, &rst); err != nil {
		return err
	}
	if rst.GetInt("errorcode") != 0 {
		return errors.New(rst.GetString("errmsg"))
	}

	return nil
}

func MenuGet(appid, secret string) (string, error) {
	accessToken := CGIToken(appid, secret).Get()

	menuGetrl := fmt.Sprintf(menu_get_url, accessToken)
	buf, err := gooHttp.NewRequest().JsonContentType().Get(menuGetrl)
	if err != nil {
		return "", err
	}

	rst := goo.Params{}
	if err := json.Unmarshal(buf, &rst); err != nil {
		return "", err
	}
	if rst.GetInt("errorcode") != 0 {
		return "", errors.New(rst.GetString("errmsg"))
	}

	return string(buf), nil
}

func MenuDelete(appid, secret string) error {
	accessToken := CGIToken(appid, secret).Get()

	menuDeleteUrl := fmt.Sprintf(menu_del_url, accessToken)
	buf, err := gooHttp.NewRequest().JsonContentType().Post(menuDeleteUrl, nil)
	if err != nil {
		return err
	}

	rst := goo.Params{}
	if err := json.Unmarshal(buf, &rst); err != nil {
		return err
	}
	if rst.GetInt("errorcode") != 0 {
		return errors.New(rst.GetString("errmsg"))
	}
	return nil
}
