package storage

import (
	"code.qingpai365.com/erp/goutil/log"
	"code.qingpai365.com/erp/goutil/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/url"
	"strings"
	"time"
)

var (
	_storageType     string
	_endpoint        string
	_accessKeyID     string
	_secretAccessKey string
	_bucketName      string
	_prefix          string
)

func Init(storageType string, endpoint string, accessKeyID string, secretAccessKey string, bucketName string, prefix string) {
	_storageType = storageType
	_endpoint = endpoint
	_accessKeyID = accessKeyID
	_secretAccessKey = secretAccessKey
	_bucketName = bucketName
	_prefix = prefix
}

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

// PutObject 上传文件
func PutObject(key string, fileHeader *multipart.FileHeader) (gin.H, error) {
	var provider Provider

	switch _storageType {
	case "aliyunoss":
		provider = NewAliyunOss(_endpoint, _accessKeyID, _secretAccessKey, _bucketName)
	case "minio":
		provider = NewMinio(_endpoint, _accessKeyID, _secretAccessKey, _bucketName)
	default:
		return nil, errors.New("invalid storage type")
	}

	if _prefix != "" {
		key = fmt.Sprintf("%s/%s", _prefix, key)
	}

	uploadInfo, err := provider.PutObject(key, fileHeader)

	if err != nil {
		log.Errorf("上传失败: er = %v", err)
		log.Errorf("上传失败: uploadInfo = %s", util.ToString(uploadInfo))
		return nil, err
	}

	duration := time.Duration(1) * time.Hour
	signedUrl, err := PresignedGet(key, duration)
	if err != nil {
		return nil, err
	}

	result := gin.H{
		"key":      key,
		"name":     fileHeader.Filename,
		"size":     fileHeader.Size,
		"url":      signedUrl,
		"expireAt": time.Now().Add(duration),
	}

	return result, nil
}

// PresignedGet 获取签名后的文件url
func PresignedGet(key string, duration time.Duration) (string, error) {
	var provider Provider

	switch _storageType {
	case "aliyunoss":
		provider = NewAliyunOss(_endpoint, _accessKeyID, _secretAccessKey, _bucketName)
	case "minio":
		provider = NewMinio(_endpoint, _accessKeyID, _secretAccessKey, _bucketName)
	default:
		return "", errors.New("invalid storage type")
	}

	urlResp, err := provider.SignUrl(key, duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("storage%s?%s", urlResp.Path, urlResp.RawQuery), nil
}

// BatchPresignedGet 批量获取签名后的文件url
func BatchPresignedGet(key string, duration time.Duration) (map[string]gin.H, error) {
	var provider Provider

	switch _storageType {
	case "aliyunoss":
		provider = NewAliyunOss(_endpoint, _accessKeyID, _secretAccessKey, _bucketName)
	case "minio":
		provider = NewMinio(_endpoint, _accessKeyID, _secretAccessKey, _bucketName)
	default:
		return nil, errors.New("invalid storage type")
	}

	keyList := strings.Split(key, ",")

	urls := make(map[string]gin.H)
	for _, key := range keyList {
		if key == "" {
			continue
		}
		signedUrl, err := provider.SignUrl(key, duration)
		if err != nil {
			return nil, err
		}
		urlString := fmt.Sprintf("storage%s?%s", signedUrl.Path, signedUrl.RawQuery)
		urls[key] = gin.H{"url": urlString, "expireAt": time.Now().Add(duration)}
	}

	return urls, nil
}
