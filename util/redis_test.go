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

func TestGetRedisObjKey(t *testing.T) {
	type args struct {
		prefix string
		param  interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal case",
			args: args{
				prefix: "military",
				param: map[string]interface{}{
					"id":   1212,
					"name": "reza",
				},
			},
			want: "military:f1cc8d38585e6066559266e9b601bc64",
		},
		{
			name: "Normal case",
			args: args{
				prefix: "military",
				param: map[string]interface{}{
					"name": "reza",
					"id":   1212,
				},
			},
			want: "military:93b5aad38d82963d41a14eea552188db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisObjKey(tt.args.prefix, tt.args.param); got != tt.want {
				t.Errorf("GetRedisObjKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
