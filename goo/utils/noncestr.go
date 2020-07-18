package gooUtils

import (
	"time"
	"strings"
)

func NonceStr() string {
	return strings.ToLower(Id2Code(time.Now().UnixNano()))
}
