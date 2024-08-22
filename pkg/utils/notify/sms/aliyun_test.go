package sms

import (
    "testing"
)

func TestSendVerifyCodeByAliyun(t *testing.T) {
    if true {
        t.Skip()
        return
    }

    cfg := &config.SMSConfig{
        Aliyun: &config.AliConfig{
            AccessKeyID:     "",
            AccessKeySecret: "",
        },
        TemplateCode: "",
        SignName:     "gin-init",
    }
    err := sendVerifyCodeByAliyun(cfg, "", "321456")
    if err != nil {
        t.Error(err)
    }

}
