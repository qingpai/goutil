package storage

import (
	"fmt"
	"github.com/spf13/viper"
	"mime/multipart"
	"net/url"
	"time"
)

type Provider interface {
	PutObject(key string, fileHeader *multipart.FileHeader) (*UploadInfo, error)
	GetObject(key string) ([]byte, error)
	SignUrl(key string, duration time.Duration) (*url.URL, error)
}

type UploadInfo struct {
	Bucket   string `json:"bucket"`
	Key      string `json:"key"`
	FileName string `json:"fileName"`
}

func PresignedGet(conf *viper.Viper, key string, duration time.Duration) (string, error) {
	storageType := conf.GetString("storage.type")
	endpoint := conf.GetString("storage.endpoint")
	accessKeyID := conf.GetString("storage.accessKeyID")
	secretAccessKey := conf.GetString("storage.secretAccessKey")
	bucketName := conf.GetString("storage.bucketName")

	var provider Provider

	switch storageType {
	case "aliyunoss":
		provider = NewAliyunOss(endpoint, accessKeyID, secretAccessKey, bucketName)
	case "minio":
		provider = NewMinio(endpoint, accessKeyID, secretAccessKey, bucketName)
	}

	urlResp, err := provider.SignUrl(key, duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/storage%s?%s", conf.GetString("server.host"), urlResp.Path, urlResp.RawQuery), nil
}
