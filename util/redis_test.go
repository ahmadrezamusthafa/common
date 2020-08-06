package util

import "testing"

func TestGetRedisKey(t *testing.T) {
	type args struct {
		prefix string
		param  map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal case",
			args: args{
				prefix: "segmentation:filterOperator",
				param: map[string]interface{}{
					"id": "xxx-abc-def-001",
				},
			},
			want: "cc4a1408126745602aa7c5368a61f022",
		},
		{
			name: "Normal case",
			args: args{
				prefix: "segmentation:filterOperator",
				param: map[string]interface{}{
					"id":    "xxx-abc-def-001",
					"email": "test01@mgal.com",
				},
			},
			want: "84be9db82381df1ec4eac52093624f9b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisKey(tt.args.prefix, tt.args.param); got != tt.want {
				t.Errorf("GetRedisKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
