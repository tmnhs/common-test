package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tmnhs/common-test/internal/model"
	"github.com/tmnhs/common/logger"
	"gorm.io/gorm"
	"os"
)

func RegisterTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("register table failed, error:%s", err.Error()))
		os.Exit(0)
	}
	entities := []model.User{
		{UserName: "root", Password: "e10adc3949ba59abbe56e057f20f883e", Role: model.RoleAdmin, Email: "333333333@qq.com"},
	}
	if exist := checkDataExist(db); !exist {
		if err := db.Table(model.CommonUserTableName).Create(&entities).Error; err != nil {
			return errors.Wrap(err, "Failed to initialize table data")
		}
	}
	logger.GetLogger().Info("register table success")
	return nil
}

func checkDataExist(db *gorm.DB) bool {
	if errors.Is(db.Table(model.CommonUserTableName).Where("username = ?", "root").First(&model.User{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
