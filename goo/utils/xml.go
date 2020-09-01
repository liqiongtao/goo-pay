package gooUtils

import (
	"bytes"
	"encoding/xml"
	"googo.io/goo"
	gooLog "googo.io/goo/log"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type StringMap map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StringMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func Xml2Map(buf []byte) goo.Params {
	var (
		key   string
		value string
		data  = goo.Params{}
	)

	decoder := xml.NewDecoder(bytes.NewReader(buf))

	for {
		decoder.Skip()
		t, err := decoder.Token()
		if err != nil {
			gooLog.Error(err.Error())
			break
		}
		switch token := t.(type) {
		case xml.StartElement:
			key = strings.TrimSpace(token.Name.Local)
			value = ""
		case xml.CharData:
			value = string([]byte(token))
		case xml.EndElement:
			continue
		}
		if key == "xml" || key == "root" || value == "\n" {
			continue
		}
		data.Set(key, strings.TrimSpace(value))
	}

	return data
}

func Map2Xml(params goo.Params) []byte {
	var buf bytes.Buffer
	buf.WriteString("<xml>")
	buf.Write(params2xml(params))
	buf.WriteString("</xml>")
	return buf.Bytes()
}

func params2xml(params goo.Params) []byte {
	var buf bytes.Buffer
	for k, v := range params {
		buf.WriteString("<")
		buf.WriteString(k)
		switch reflect.TypeOf(v).String() {
		case "string":
			buf.WriteString("><![CDATA[")
			buf.WriteString(v.(string))
			buf.WriteString("]]></")
		case "int":
			buf.WriteString("><![CDATA[")
			buf.WriteString(strconv.FormatInt(int64(v.(int)), 10))
			buf.WriteString("]]></")
		case "int64":
			buf.WriteString("><![CDATA[")
			buf.WriteString(strconv.FormatInt(v.(int64), 10))
			buf.WriteString("]]></")
		case "bool":
			buf.WriteString("><![CDATA[")
			buf.WriteString(strconv.FormatBool(v.(bool)))
			buf.WriteString("]]></")
		case "float64":
			buf.WriteString("><![CDATA[")
			buf.WriteString(strconv.FormatFloat(v.(float64), 'f', -1, 64))
			buf.WriteString("]]></")
		case "goo.Params":
			buf.WriteString(">")
			buf.Write(params2xml(v.(goo.Params)))
			buf.WriteString("</")
		}
		buf.WriteString(k)
		buf.WriteString(">")
	}
	return buf.Bytes()
}
