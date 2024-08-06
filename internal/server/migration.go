package server

import (
    "context"
    "gin-init/internal/model/model_type"
    "gin-init/pkg/log"
    "os"

    "go.uber.org/zap"
    "gorm.io/gorm"
)

type Migrate struct {
    db  *gorm.DB
    log *log.Logger
}

func NewMigrate(db *gorm.DB, log *log.Logger) *Migrate {
    return &Migrate{
        db:  db,
        log: log,
    }
}
func (m *Migrate) Start(ctx context.Context) error {
    if err := m.db.AutoMigrate(&model_type.User{}); err != nil {
        m.log.Error("user migrate error", zap.Error(err))
        return err
    }
    m.log.Info("AutoMigrate success")
    os.Exit(0)
    return nil
}
func (m *Migrate) Stop(ctx context.Context) error {
    m.log.Info("AutoMigrate stop")
    return nil
}
