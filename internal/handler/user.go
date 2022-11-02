package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tmnhs/common"
	"github.com/tmnhs/common-test/internal/middlerware"
	"github.com/tmnhs/common-test/internal/model"
	"github.com/tmnhs/common-test/internal/model/request"
	"github.com/tmnhs/common-test/internal/model/resp"
	"github.com/tmnhs/common-test/internal/service"
	"github.com/tmnhs/common/logger"
	"github.com/tmnhs/common/utils"
	"time"
)

type UserRouter struct {
}

var defaultUserRouter = new(UserRouter)

func (u *UserRouter) Login(c *gin.Context) {
	var req request.ReqUserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[user_login] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[user_login] request parameter error", c)
		return
	}
	user, err := service.DefaultUserService.Login(req.UserName, req.Password)
	if err != nil || user.ID == 0 {
		logger.GetLogger().Error(fmt.Sprintf("[user_login] db error:%v", err))
		common.FailWithMessage(common.ERROR, "[user_login] username or password is incorrect", c)
		return
	}

	j := middlerware.NewJWT()
	claims := j.CreateClaims(middlerware.BaseClaims{
		ID:       user.ID,
		UserName: user.UserName,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		common.FailWithMessage(common.ErrorTokenGenerate, "获取token失败", c)
		return
	}
	common.OkWithDetailed(resp.RspLogin{
		User:  user,
		Token: token,
	}, "login success", c)
}

func (u *UserRouter) Register(c *gin.Context) {
	var req request.ReqUserRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[user_register] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[user_register] request parameter error", c)
		return
	}
	user, err := service.DefaultUserService.FindByUserName(req.UserName)
	if user != nil || err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[user_register] db find by username:%s", req.UserName))
		common.FailWithMessage(common.ErrorUserNameExist, "[user_register] the user name has already been used", c)
		return
	}
	if req.Role == 0 {
		req.Role = model.RoleNormal
	}
	userModel := &model.User{
		UserName: req.UserName,
		Password: utils.MD5(req.Password),
		Role:     req.Role,
		Email:    req.Email,
		Created:  time.Now().Unix(),
	}
	insertId, err := userModel.Insert()
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[user_register] db insert error:%v", err))
		common.FailWithMessage(common.ERROR, "[user_register] db insert error", c)
		return
	}
	userModel.ID = insertId
	common.OkWithDetailed(userModel, "register success", c)
}

func (u *UserRouter) Update(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[update_user] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[update_user] request parameter error", c)
		return
	}
	req.Updated = time.Now().Unix()
	err := req.Update()
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[update_user] db update error:%v", err))
		common.FailWithMessage(common.ERROR, "[update_user] db update error", c)
		return
	}
	common.OkWithMessage("update success", c)
}

func (u *UserRouter) Delete(c *gin.Context) {
	var req common.ByIDS
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[delete_user] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[delete_user] request parameter error", c)
		return
	}
	for _, id := range req.IDS {
		userModel := model.User{ID: id}
		err := userModel.Delete()
		if err != nil {
			logger.GetLogger().Error(fmt.Sprintf("[delete_user] db error:%v", err))
		}
	}
	common.OkWithMessage("delete success", c)
}

func (u *UserRouter) ChangePassword(c *gin.Context) {
	var req request.ReqChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[change_password] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[change_password] request parameter error", c)
		return
	}
	err := service.DefaultUserService.ChangePassword(middlerware.GetUserInfo(c).ID, req.Password, req.NewPassword)
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[change_password] db error:%v", err))
		common.FailWithMessage(common.ERROR, "[change_password] db error", c)
		return
	}
	common.OkWithMessage("update success", c)
}

func (u *UserRouter) FindById(c *gin.Context) {
	var req common.ByID
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[find_user] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[find_user] request parameter error", c)
		return
	}
	if req.ID == 0 {
		req.ID = middlerware.GetUserInfo(c).ID
	}
	userModel := model.User{ID: req.ID}
	err := userModel.FindById()
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[find_user] db  error:%v", err))
		common.FailWithMessage(common.ERROR, "[find_user] db  error", c)
		return
	}
	common.OkWithDetailed(userModel, "find success", c)
}

func (u *UserRouter) Search(c *gin.Context) {
	var req request.ReqUserSearch
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[search_user] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[search_user] request parameter error", c)
		return
	}
	req.Check()
	users, total, err := service.DefaultUserService.Search(&req)
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[search_user] db error:%v", err))
		common.FailWithMessage(common.ERROR, "[search_user] db error", c)
		return
	}
	common.OkWithDetailed(common.PageResult{
		List:     users,
		Total:    total,
		PageSize: req.PageSize,
		Page:     req.Page,
	}, "search success", c)
}
