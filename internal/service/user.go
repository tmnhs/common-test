package service

import (
	"github.com/tmnhs/common-test/internal/model"
	"github.com/tmnhs/common-test/internal/model/request"
	"github.com/tmnhs/common/dbclient"
	"github.com/tmnhs/common/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

var DefaultUserService = new(UserService)

func (us *UserService) Login(username, password string) (u *model.User, err error) {
	u = new(model.User)
	//返回没有密码的用户信息
	err = dbclient.GetMysqlDB().Select("id", "username", "email", "role", "created", "updated").Table(model.CommonUserTableName).Where("username = ? And password = ?", username, utils.MD5(password)).Find(u).Error
	return
}

func (us *UserService) FindByUserName(username string) (u *model.User, err error) {
	u = new(model.User)
	err = dbclient.GetMysqlDB().Table(model.CommonUserTableName).Where("username = ? ", username).First(u).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (us *UserService) ChangePassword(userId int, oldPassword, newPassword string) error {
	return dbclient.GetMysqlDB().Table(model.CommonUserTableName).Where("id = ? And password =? ", userId, utils.MD5(oldPassword)).Update("password", utils.MD5(newPassword)).Error
}

func (us *UserService) Search(s *request.ReqUserSearch) ([]model.User, int64, error) {
	db := dbclient.GetMysqlDB().Table(model.CommonUserTableName)
	if len(s.UserName) > 0 {
		db = db.Where("username like ?", s.UserName+"%")
	}

	if len(s.Email) > 0 {
		db.Where("email = ?", s.Email)
	}
	if s.Role > 0 {
		db.Where("role = ?", s.Role)
	}
	if s.ID > 0 {
		db.Where("id = ?", s.ID)
	}
	users := make([]model.User, 2)
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//不建议直接搜索密码
	err = db.Select("id", "username", "email", "role", "created", "updated").Limit(s.PageSize).Offset((s.Page - 1) * s.PageSize).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
