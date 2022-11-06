package service

import (
	"github.com/tmnhs/common"
	"github.com/tmnhs/common-test/internal/model"
	"mime/multipart"
)

type UploadService struct {
}

var DefaultUploadService = new(UploadService)

func (u *UploadService) UploadFile(header *multipart.FileHeader) (upload *model.Upload, err error) {
	upload = new(model.Upload)
	switch common.GetConfigModels().System.UploadType {
	case "local":
		upload.Url, upload.Key, err = common.GetConfigModels().Upload.Local.UploadFile(header)
	case "qiniu":
		upload.Url, upload.Key, err = common.GetConfigModels().Upload.Qiniu.UploadFile(header)
	case "tencent-cos":
		upload.Url, upload.Key, err = common.GetConfigModels().Upload.TencentCOS.UploadFile(header)
	case "aliyun-oss":
		upload.Url, upload.Key, err = common.GetConfigModels().Upload.AliyunOSS.UploadFile(header)
	case "huawei-obs":
		upload.Url, upload.Key, err = common.GetConfigModels().Upload.HuaWeiObs.UploadFile(header)
	default:
		upload.Url, upload.Key, err = common.GetConfigModels().Upload.Local.UploadFile(header)
	}
	return
}

func (u *UploadService) DeleteFile(key string) (err error) {
	switch common.GetConfigModels().System.UploadType {
	case "local":
		err = common.GetConfigModels().Upload.Local.DeleteFile(key)
	case "qiniu":
		err = common.GetConfigModels().Upload.Qiniu.DeleteFile(key)
	case "tencent-cos":
		err = common.GetConfigModels().Upload.TencentCOS.DeleteFile(key)
	case "aliyun-oss":
		err = common.GetConfigModels().Upload.AliyunOSS.DeleteFile(key)
	case "huawei-obs":
		err = common.GetConfigModels().Upload.HuaWeiObs.DeleteFile(key)
	default:
		err = common.GetConfigModels().Upload.Local.DeleteFile(key)
	}
	return
}
