package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

func GetRedisKey(prefix string, param map[string]interface{}) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(prefix)

	keys := make([]string, 0, len(param))
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		key := strings.Replace(key, " ", "", -1)
		val := strings.Replace(CastToString(param[key]), " ", "", -1)
		buffer.WriteByte(':')
		buffer.WriteString(key)
		buffer.WriteByte(':')
		buffer.WriteString(val)
	}
	return hashRedisKey(buffer.String())
}

func hashRedisKey(key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
