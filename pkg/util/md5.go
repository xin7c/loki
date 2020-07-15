package util

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	md5String := hex.EncodeToString(m.Sum(nil))
	log.Printf("[%s] md5 ==> [%s]", value, md5String)
	return md5String
}
