package encryptor

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"hash"
)

func MD5(v string) string {
	h := md5.New()
	h.Write([]byte(v))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacMD5(key, str string) string {
	return HmacHash(md5.New, key, str)
}

func HmacHash(h func() hash.Hash, key, str string) string {
	mac := hmac.New(h, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}
