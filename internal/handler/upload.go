package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tmnhs/common"
	"github.com/tmnhs/common-test/internal/model"
	"github.com/tmnhs/common-test/internal/service"
	"github.com/tmnhs/common/logger"
)

type UploadRouter struct {
}

var defaultUploadRouter = new(UploadRouter)

func (u *UploadRouter) UploadFile(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[upload_file] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[upload_file] request parameter error", c)
		return
	}
	upload, err := service.DefaultUploadService.UploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[upload_file] error:%s", err.Error()))
		common.FailWithMessage(common.ERROR, fmt.Sprintf("[upload_file] error:%s", err.Error()), c)
		return
	}
	common.OkWithDetailed(upload, "upload success", c)
}

func (u *UploadRouter) DeleteFile(c *gin.Context) {
	var req model.Upload
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[delete_file] request parameter error:%s", err.Error()))
		common.FailWithMessage(common.ErrorRequestParameter, "[delete_file] request parameter error", c)
		return
	}
	if err := service.DefaultUploadService.DeleteFile(req.Key); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("[delete_file] error:%s", err.Error()))
		common.FailWithMessage(common.ERROR, fmt.Sprintf("[delete_file] error:%s", err.Error()), c)
		return
	}
	common.OkWithMessage("delete success", c)
}
