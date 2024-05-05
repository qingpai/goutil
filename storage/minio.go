package storage

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

type Minio struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName      string `json:"bucketName"`
}

// NewMinio new minio
func NewMinio(endpoint, accessKeyId, accessKeySecret, bucketName string) *Minio {
	return &Minio{
		Endpoint:        endpoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		BucketName:      bucketName,
	}
}

func (m *Minio) SignUrl(key string, duration time.Duration) (*url.URL, error) {
	client, err := minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKeyId, m.AccessKeySecret, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	reqParams := make(url.Values)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return client.PresignedGetObject(ctx, m.BucketName, key, duration, reqParams)
}

func (m *Minio) GetObject(key string) ([]byte, error) {
	client, err := minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKeyId, m.AccessKeySecret, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	file, err := client.GetObject(ctx, m.BucketName, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	_, err = file.Stat()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)

	return buf.Bytes(), nil
}

func (m *Minio) PutObject(key string, fileHeader *multipart.FileHeader) (*UploadInfo, error) {
	fileExt := strings.ToLower(filepath.Ext(fileHeader.Filename))
	client, err := minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKeyId, m.AccessKeySecret, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	file, err := fileHeader.Open()
	defer file.Close()

	_, err = client.PutObject(ctx, m.BucketName, key, file, fileHeader.Size, minio.PutObjectOptions{ContentType: ExtToContentType(fileExt)})
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

func ExtToContentType(ext string) string {
	switch ext {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".png":
		return "image/png"
	case ".jpg":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".ico":
		return "image/x-icon"
	case ".txt":
		return "text/plain"
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".pdf":
		return "application/pdf"
	case ".zip":
		return "application/zip"
	case ".mp3":
		return "audio/mpeg"
	case ".mp4":
		return "video/mp4"
	case ".wav":
		return "audio/wav"
	case ".ogg":
		return "audio/ogg"
	case ".webm":
		return "video/webm"
	case ".wasm":
		return "application/wasm"
	default:
		return "application/octet-stream"
	}
}
