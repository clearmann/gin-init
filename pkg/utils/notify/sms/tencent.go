package sms

// TencentConfig 腾讯云配置
type TencentConfig struct {
    SecretID  string `yaml:"secret_id"`
    SecretKey string `yaml:"secret_key"`
    Region    string `yaml:"region"`
    Endpoint  string `yaml:"endpoint"`
}
