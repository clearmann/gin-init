package repository

import (
    "encoding/json"
    "gin-init/internal/model/model_type"
    "gin-init/internal/model/model_type/settings"

    "gopkg.in/yaml.v3"
    "gorm.io/gorm/clause"
)

type SettingRepository interface {
    GetByID(id uint) (*settings.CoreSetting, error)
    UpdateSettings(id uint) (*settings.CoreSetting, error)
    CreateSettings(v *settings.CoreSetting) error
    Get(group, key string) (*settings.CoreSetting, error)
    GetYaml(group, key string, value interface{}) error
    GetText(group, key string) (string, error)
    GetPassword(group, key string) (string, error)
    GetJSON(group, key string, value interface{}) error
    Updates(sets ...*settings.CoreSetting) error
    List(group string, keys ...string) ([]*settings.CoreSetting, error)
}

func NewSettingRepository(r *Repository) SettingRepository {
    return &settingRepository{Repository: r}
}

type settingRepository struct {
    *Repository
}

// GetByID .
func (r *settingRepository) GetByID(id uint) (*settings.CoreSetting, error) {
    item := new(settings.CoreSetting)
    err := r.db.
        Where("id = ?", id).
        Find(item).Error
    if err != nil {
        return nil, err
    }
    return item, nil
}

func (r *settingRepository) UpdateSettings(id uint) (*settings.CoreSetting, error) {
    item := new(settings.CoreSetting)
    err := r.db.
        Where("id = ?", id).
        Find(item).Error
    if err != nil {
        return nil, err
    }
    return item, nil
}

func (r *settingRepository) CreateSettings(v *settings.CoreSetting) error {
    return r.db.
        Clauses(clause.OnConflict{
            DoUpdates: clause.AssignmentColumns([]string{"name", "describe", "value", "value_type", "default"}),
        }).Create(v).Error
}

// Get .
func (r *settingRepository) Get(group, key string) (*settings.CoreSetting, error) {
    item := new(settings.CoreSetting)
    err := r.db.
        Where("`group` = ? AND `key` = ?", group, key).
        First(item).Error
    if err != nil {
        return nil, err
    }
    return item, nil
}

// GetYaml 获取yaml配置
func (r *settingRepository) GetYaml(group, key string, value interface{}) error {
    si, err := r.Get(group, key)
    if err != nil {
        return err
    }
    return yaml.Unmarshal([]byte(si.Value), value)
}

// GetText 获取文本配置
func (r *settingRepository) GetText(group, key string) (string, error) {
    si, err := r.Get(group, key)
    if err != nil {
        return "", err
    }
    return si.Value, nil
}

// GetPassword 获取密码配置
func (r *settingRepository) GetPassword(group, key string) (string, error) {
    si, err := r.Get(group, key)
    if err != nil {
        return "", err
    }
    return settings.DecryptPassword(si.Value), nil
}

// GetJSON 获取json配置
func (r *settingRepository) GetJSON(group, key string, value interface{}) error {
    si, err := r.Get(group, key)
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(si.Value), value)
}

// List 配置列表
func (r *settingRepository) List(group string, keys ...string) ([]*settings.CoreSetting, error) {
    var ret []*settings.CoreSetting
    err := r.db.Table(model_type.TableNameCoreSetting).
        Where("`group` = ? AND `key` IN (?)", group, keys).
        Find(&ret).Error
    if err != nil {
        return nil, err
    }
    return ret, nil
}

// Updates 更新settings值
func (r *settingRepository) Updates(sets ...*settings.CoreSetting) error {
    for _, set := range sets {
        sql := r.db.Table(model_type.TableNameCoreSetting)
        if set.ID != 0 {
            sql = sql.Where("id = ?", set.ID)
        } else {
            sql = sql.Where("`group` = ? AND `key` = ?", set.Group, set.Key)
        }
        if set.ValueType == settings.ValuePassword {
            set.Value = settings.EncryptPassword(set.Value)
        }
        err := sql.Update("value", set.Value).Error
        if err != nil {
            r.logger.Error("[settings] update %s failed, %s")
            return err
        }
    }
    return nil
}
