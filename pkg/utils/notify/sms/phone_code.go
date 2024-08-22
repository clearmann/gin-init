package sms

import (
    "context"
    "fmt"
)

type SMSConfig struct {
    Aliyun       *AliConfig     `yaml:"aliyun"`
    Tencent      *TencentConfig `yaml:"tencent"`
    SignName     string         `yaml:"sign_name"`
    TemplateCode string         `yaml:"template_code"`
}

// SendPhoneCode 发送验证码
func SendPhoneCode(ctx context.Context, phone, code string, cfg *SMSConfig) error {
    var err error
    if cfg.Aliyun != nil {
        err = sendVerifyCodeByAliyun(cfg, phone, code)
        if err != nil {
        }
    } else if cfg.Tencent != nil {
        // TODO 腾讯云发送短信 未配置
    } else {
        err = fmt.Errorf("sms config is empty")
    }
    return err
}
