package util

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"unicode/utf8"
)

// VerifyMobile 验证手机号
func VerifyMobile(mobile string) bool {
	regular := "^((13[0-9])|(14[0-9])|(15[0-9])|(16[0-9])|(17[0-9])|(18[0-9])|(19[0-9]))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

// VerifySmsContent	验证短信内容
func VerifySmsContent(content string) error {
	clen := utf8.RuneCountInString(content)
	if clen < 1 {
		return errors.New("短信内容不能为空")
	} else if clen > 500 {
		return errors.New("短信内容不能超过500字")
	}

	return nil
}

// GenSmsContent 生成短信内容
func GenSmsContent(smsType int, sign string, content string, rejectText string) string {
	return fmt.Sprintf("【%s】%s%s", sign, content, rejectText)
}

// SmsCalcBillingCount
//
//	@Description: 计算短信条数, <70字时，按1条计算，大于70字时，按67字/条计算
//	@param smsContent
//	@return int
func SmsCalcBillingCount(smsContent string) int {
	contentLength := utf8.RuneCountInString(smsContent)
	smsCount := 1
	if contentLength > 70 {
		// ! 70 字及以内为一条，超出 70 字后所有内容按 67 字/条进行切割计数。
		smsCount = int(math.Ceil(float64(contentLength) / float64(67)))
	}
	return smsCount
}
