package validator

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var (
	uni *ut.UniversalTranslator
	//validate *validator.Validate
	trans ut.Translator
)

func Init() {
	//注册翻译器
	zhInstance := zh.New()
	uni = ut.New(zhInstance, zhInstance)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("display")
	})
	//注册翻译器
	err := zhtranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

// Translate 翻译错误信息
func Translate(err error) string {
	var result string

	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		for _, e := range validationError {
			errMessage := e.Translate(trans)
			result += errMessage + ";"
		}
		return result[:len(result)-1]
	}

	return err.Error()
}
