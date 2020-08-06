package util

import "testing"

func BenchmarkGetRedisKey(b *testing.B) {
	prefix := "segmentation:filterOperator"
	param := map[string]interface{}{
		"id":    "xxx-abc-def-001",
		"email": "test01@mgal.com",
	}
	for n := 0; n < b.N; n++ {
		GetRedisKey(prefix, param)
	}
}

func BenchmarkGetRedisObjKey(b *testing.B) {
	prefix := "segmentation:filterOperator"
	param := map[string]interface{}{
		"id":    "xxx-abc-def-001",
		"email": "test01@mgal.com",
	}
	for n := 0; n < b.N; n++ {
		GetRedisObjKey(prefix, param)
	}
}
