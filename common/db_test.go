package common

import "testing"

func TestConnRedis(t *testing.T) {
	InitRedis()
}
func TestSetValue(t *testing.T) {
	InitRedis()
	Redis.Set("key", "value", 0)
}
func TestGetValue(t *testing.T) {
	InitRedis()
	Redis.Get("key")
}

func TestConnSql(t *testing.T) {
	InitMysql()
}
