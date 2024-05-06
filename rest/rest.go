package rest

import (
	"code.qingpai365.com/erp/goutil/log"
	"errors"
	"github.com/go-resty/resty/v2"
	"time"
)

var client *resty.Client

func Init(retryCount int, proxyUrl string, debug bool) {
	client = resty.New()

	if debug {
		client.Debug = true
		client.EnableTrace()
	}

	client.SetRetryCount(retryCount)
	client.SetRetryWaitTime(time.Duration(3) * time.Second)
	client.SetRetryMaxWaitTime(time.Duration(5) * time.Second)
	if proxyUrl != "" {
		client.SetProxy(proxyUrl)
	}
	client.OnError(func(req *resty.Request, err error) {
		var v *resty.ResponseError
		if errors.As(err, &v) {
			log.Errorf("resty response: %v", v.Response)
			log.Errorf("resty original error: %v", v.Err)
		}
		log.Errorf("resty error: %v", err)
	})
}

func GetClient() *resty.Client {
	return client
}
