package constants

import "fmt"

const (
	RedisKeyChatMessage      = "subscribe:aigc:chat:question"
	RedisKeyChatMessageQueue = "queue:aigc:chat:question"
)

func RedisKeyGithubOauthUserInfo(githubid uint) string {
	return fmt.Sprintf("github:oauth:user:%d", githubid)
}

func RedisKeyWorkWechatOauthUserInfo(userid string) string {
	return fmt.Sprintf("workwechat:oauth:user:%s", userid)
}

func RedisKeyWechatWebOauthUserInfo(unionid string) string {
	return fmt.Sprintf("wechat_web:oauth:user:%s", unionid)
}
func RedisKeyPhoneVerifyCode(phone string) string {
	return fmt.Sprintf("phone_verify_code:phone:%s", phone)
}
