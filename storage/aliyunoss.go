package storage

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"mime/multipart"
	"net/url"
	"time"
)

type AliyunOss struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName      string `json:"bucketName"`
}

// NewAliyunOss new aliyun oss
func NewAliyunOss(endpoint, accessKeyId, accessKeySecret, bucketName string) *AliyunOss {
	return &AliyunOss{
		Endpoint:        endpoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		BucketName:      bucketName,
	}
}

func (m *AliyunOss) SignUrl(key string, duration time.Duration) (*url.URL, error) {
	client, err := oss.New(m.Endpoint, m.AccessKeyId, m.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(m.BucketName)
	if err != nil {
		return nil, err
	}

	signedUrl, err := bucket.SignURL(key, oss.HTTPGet, int64(duration))
	if err != nil {
		return nil, err
	}

	return url.Parse(signedUrl)
}

func (m *AliyunOss) GetObject(key string) ([]byte, error) {
	client, err := oss.New(m.Endpoint, m.AccessKeyId, m.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(m.BucketName)
	if err != nil {
		return nil, err
	}

	body, err := bucket.GetObject(key)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(body)
}

func (m *AliyunOss) PutObject(key string, fileHeader *multipart.FileHeader) (*UploadInfo, error) {
	client, err := oss.New(m.Endpoint, m.AccessKeyId, m.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(m.BucketName)
	if err != nil {
		return nil, err
	}

	file, err := fileHeader.Open()
	defer file.Close()

	err = bucket.PutObject(key, file)
	if err != nil {
		return nil, err
	}

	uploadInfo := &UploadInfo{
		Bucket:   m.BucketName,
		Key:      key,
		FileName: fileHeader.Filename,
	}

	return uploadInfo, nil
}
