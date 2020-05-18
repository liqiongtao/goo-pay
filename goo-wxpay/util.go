package gooWXPay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Success() string {
	return `<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>`
}

func obj2xml(stc interface{}) []byte {
	var buf bytes.Buffer

	buf.WriteString("<xml>")

	fields := reflect.TypeOf(stc).Elem()
	values := reflect.ValueOf(stc).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i).Tag.Get("xml")
		value := fmt.Sprint(values.Field(i).Interface())
		if value == "" {
			continue
		}
		buf.WriteString(fmt.Sprintf("<%s>%s</%s>", field, value, field))
	}

	buf.WriteString("</xml>")

	return buf.Bytes()
}

func obj2querystring(stc interface{}) string {
	params := []string{}

	fields := reflect.TypeOf(stc).Elem()
	values := reflect.ValueOf(stc).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i).Tag.Get("xml")
		value := fmt.Sprint(values.Field(i).Interface())
		if value == "" {
			continue
		}
		params = append(params, fmt.Sprintf("%s=%s", field, value))
	}

	sort.Strings(params)

	return strings.Join(params, "&")
}

func map2querystring(data map[string]interface{}) string {
	params := []string{}

	for key, value := range data {
		if key == "sign" || value == "" {
			continue
		}
		params = append(params, fmt.Sprintf("%s=%s", key, value))
	}

	sort.Strings(params)

	return strings.Join(params, "&")
}

func xml2map(buf []byte) map[string]interface{} {
	field := ""
	params := map[string]interface{}{}

	decoder := xml.NewDecoder(bytes.NewBuffer(buf))
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			field = token.Name.Local
			if field == "xml" {
				continue
			}

		case xml.CharData:
			value := strings.TrimSpace(string([]byte(token)))
			if field == "" || value == "" {
				continue
			}
			params[field] = value
			field = ""
		}
	}

	return params
}
