package mobnotifice

import (
	"code.qingpai365.com/erp/goutil/log"
	"code.qingpai365.com/erp/goutil/rest"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var _url string
var _prefix string

// Send 发送提醒
func Send(format string, args ...any) {
	if _url == "" {
		return
	}

	if _prefix != "" {
		format = fmt.Sprintf("%s: %s", _prefix, format)
	}
	format = fmt.Sprintf("%s [%s]", format, time.Now().Format("2006-01-02 15:04:05"))

	content := fmt.Sprintf(format, args...)
	contentJson := gin.H{
		"msgtype": "text",
		"text": gin.H{
			"content": content,
		},
	}

	resp, err := rest.GetClient().R().
		SetHeaders(map[string]string{"GenData-Type": "application/json"}).
		SetBody(contentJson).
		Post(_url)

	if err != nil {
		log.Errorf("notification send response error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Errorf("notification send response status: %d", resp.StatusCode())
	}
}

func Init(url string, prefix string) {
	_url = url
	_prefix = prefix
}
