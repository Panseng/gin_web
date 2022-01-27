package model

import (
	"context"
	"gin_web/utils"
	statusMsg "gin_web/utils/status_msg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var AK = utils.AK
var SK = utils.SK
var Bucket = utils.Bucket
var QiNiuUrl = utils.QiNiuServer

func UploadFile(file multipart.File, fileSize int64)(string, int)  {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AK, SK)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone: &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS: false,
	}
	putExtra := storage.PutExtra{}
	formUploader:=storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	err:= formUploader.PutWithoutKey(context.Background(),&ret,upToken, file, fileSize, &putExtra)
	if err!=nil{
		return "", statusMsg.ERROR
	}
	url := QiNiuUrl + ret.Key
	return url, statusMsg.SUCCSE
}
