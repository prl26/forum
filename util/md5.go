/**
* @Author: 云坠
* @Date: 2022/5/28 21:16
**/
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

