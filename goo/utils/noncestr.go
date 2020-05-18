package gooUtils

import (
	"strconv"
	"time"
)

func NonceStr() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
