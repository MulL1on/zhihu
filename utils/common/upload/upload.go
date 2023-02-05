package upload

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	g "juejin/app/global"
	"mime/multipart"
)

func ToQiniu(file *multipart.FileHeader, path string) (int, string) {
	ak := g.Config.Upload.AccessKey
	sk := g.Config.Upload.SecretKey
	bucket := g.Config.Upload.Bucket
	imgUrl := g.Config.Upload.Server

	src, err := file.Open()
	if err != nil {
		return 10011, err.Error()
	}
	defer src.Close()

	putPlicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(ak, sk)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	key := "image/" + path
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)
	if err != nil {
		code := 501
		return code, err.Error()
	}
	url := imgUrl + ret.Key // 返回上传后的文件访问路径
	return 0, url

}
