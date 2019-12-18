package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"loanmarket-server/lib/util"
	"loanmarket-server/thirdparty/submail/lib"
	"loanmarket-server/thirdparty/submail/sms"

	"github.com/yeeyuntech/yeego"
)

type SubmailConfig struct {
	AppId               string `json:"appId"`
	AppKey              string `json:"appKey"`
	InternationalAppId  string `json:"internationalAppId"`
	InternationalAppKey string `json:"internationalAppKey"`
}

type SubmailResponse struct {
	Status string
	Code   int64
	Msg    string
}

var (
	submailConfig = SubmailConfig{}

	enProjectId = "ZRFOQ4"
)

func Init() {
	submailConfig.AppId = yeego.Config.GetString("submail.AppId")
	submailConfig.AppKey = yeego.Config.GetString("submail.AppKey")
	submailConfig.InternationalAppId = yeego.Config.GetString("submail.InternationalAppId")
	submailConfig.InternationalAppKey = yeego.Config.GetString("submail.InternationalAppKey")
}

func SendSmsCode(to, code string) error {
	phone, _ := util.ParsePhoneNumber(to)
	var conf lib.Config
	if phone.AreaCode == "86" {
		conf = lib.NewConfig(submailConfig.AppId, submailConfig.AppKey, "md5")
	} else {
		conf = lib.NewConfig(submailConfig.InternationalAppId, submailConfig.InternationalAppKey, "md5")
	}
	submail := sms.CreateXsend(conf)
	submail.AddVar("code", code)
	submail.SetProject(enProjectId)
	if phone.AreaCode == "86" {
		submail.SetTo(phone.Phone)
	} else {
		submail.SetTo("+" + to)
	}
	ret := submail.Xsend()
	var resp SubmailResponse
	err := json.Unmarshal([]byte(ret), &resp)
	if err != nil {
		fmt.Println(ret)
		return fmt.Errorf("submail返回值解析失败:%s", err.Error())
	}
	if resp.Status != "success" {
		return errors.New(resp.Msg)
	}
	return nil
}
