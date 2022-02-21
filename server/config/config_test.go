package config

import "testing"

func TestConfig(t *testing.T) {
	if err := loadEnv(); err != nil {
		t.Errorf("err:%v", err)
	}
}
