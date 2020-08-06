package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	jsoniter "github.com/json-iterator/go"
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

func GetRedisObjKey(prefix string, param interface{}) string {
	objectStr, _ := jsoniter.MarshalToString(param)
	return prefix + ":" + hashRedisKey(objectStr)
}

func hashRedisKey(key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
