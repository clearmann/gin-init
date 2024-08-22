package sms

import (
    "fmt"

    "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliConfig struct {
    RegionID        string `yaml:"region_id"`
    Endpoint        string `yaml:"endpoint"`
    AccessKeyID     string `yaml:"access_key_id"`
    AccessKeySecret string `yaml:"access_key_secret"`
}

func sendVerifyCodeByAliyun(cfg *SMSConfig, phone string, code string) error {
    if cfg.Aliyun == nil {
        return fmt.Errorf("aliyun config is empty")
    }
    if cfg.Aliyun.AccessKeyID == "" || cfg.Aliyun.AccessKeySecret == "" {
        return fmt.Errorf("aliyun access key id or access key secret is empty")
    }
    cli, err := dysmsapi.NewClientWithAccessKey(cfg.Aliyun.RegionID, cfg.Aliyun.AccessKeyID, cfg.Aliyun.AccessKeySecret)
    if err != nil {
        return err
    }
    req := dysmsapi.CreateSendSmsRequest()

    req.SignName = cfg.SignName
    req.TemplateCode = cfg.TemplateCode
    req.PhoneNumbers = phone
    req.TemplateParam = `{"code":"` + code + `"}`
    resp, err := cli.SendSms(req)
    if err != nil {
        return err
    }
    if resp.Code != "OK" {
        return fmt.Errorf("send sms failed, code: %s, message: %s", resp.Code, resp.Message)
    }
    return nil
}
