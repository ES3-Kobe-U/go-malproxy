package service

import "testing"

func TestConfig(t *testing.T) {
	if _, err := LoadRakutenEnv(); err != nil {
		t.Errorf("err:%v", err)
	}
	if _, err := LoadAmazonEnv(); err != nil {
		t.Errorf("err:%v", err)
	}
}
