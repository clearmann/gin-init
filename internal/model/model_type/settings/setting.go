package settings

import (
    "fmt"
    "gin-init/internal/model/model_type"
    "gin-init/pkg/utils/encryptor"
    "strings"

    "gorm.io/gorm"
)

type ValueType string

const (
    ValuePassword ValueType = "password"
    ValueText     ValueType = "text"
    ValueInt64    ValueType = "int64"
    ValueBool     ValueType = "bool"
    ValueJSON     ValueType = "json"
    ValueYaml     ValueType = "yaml"
)

// CoreSetting 系统配置
type CoreSetting struct {
    gorm.Model

    Group     string    `json:"group" gorm:"column:group;type:varchar(16);uniqueIndex:idx_setting,priority:2,unique"`
    Key       string    `json:"key" gorm:"column:key;type:varchar(64);uniqueIndex:idx_setting,priority:3,unique"`
    Name      string    `json:"name" gorm:"column:name;type:varchar(64)"`
    Describe  string    `json:"describe" gorm:"column:describe;type:varchar(128)"`
    ValueType ValueType `json:"value_type" gorm:"column:value_type;default:text;type:varchar(16)"`
    Value     string    `json:"value" gorm:"column:value;type:text"`
    Default   string    `json:"default" gorm:"column:default;type:text"`
}

func (*CoreSetting) TableName() string { return model_type.TableNameCoreSetting }
func (item *CoreSetting) Identify() string {
    return fmt.Sprintf("%s/%s", item.Group, item.Key)
}

const (
    passwordAESKey          = "T651qzaEFL6Dpudy"
    passwordEncryptedPrefix = "encryped:"
)

// DecryptPassword 对加密的密码进行解密
func DecryptPassword(encData string) string {
    if !strings.HasPrefix(encData, passwordEncryptedPrefix) {
        return encData
    }
    oriData := strings.TrimPrefix(encData, passwordEncryptedPrefix)
    decrypted, err := encryptor.AesDecryptFromBase64([]byte(passwordAESKey), oriData)
    if err != nil {
        return oriData
    }
    return string(decrypted)
}

// PasswordValue 密码原文
func (item *CoreSetting) PasswordValue() string {
    if item.ValueType == ValuePassword {
        return DecryptPassword(item.Value)
    }
    return item.Value
}

// EncryptPassword 对密码进行加密
func EncryptPassword(oriData string) string {
    if strings.HasPrefix(oriData, passwordEncryptedPrefix) {
        return oriData
    }
    encrypted, err := encryptor.AesEncryptToBase64([]byte(passwordAESKey), []byte(oriData))
    if err != nil {
        return oriData
    }
    return passwordEncryptedPrefix + encrypted
}
