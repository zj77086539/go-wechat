package wechat

import (
	"fmt"
	"testing"
)

func init() {
	AppID = "wx07634c55bbebbd3b"
	AppSecret = "c5dfcacc59cb91103facf83eabd22341"
}

func TestExecute(t *testing.T) {
	// 小程序登录
	url := "https://api.weixin.qq.com/sns/jscode2session"
	param := Parameter{
		"js_code":    "登录时获取的 code",
		"grant_type": "authorization_code",
	}
	json, err := Execute(url, param)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json)
}
