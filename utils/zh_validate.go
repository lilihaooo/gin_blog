package utils

import (
	"blog_gin/pkg/res"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func ZhValidate(cr any) (validateResponse []res.ValidateResponse) {
	// 中文翻译器
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	//实例化验证器
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 注册翻译器到校验器
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
	if errs := validate.Struct(cr); errs != nil {
		if errs != nil {
			for _, e := range errs.(validator.ValidationErrors) {
				validateResponse = append(validateResponse, res.ValidateResponse{Field: strings.ToLower(e.StructField()), Msg: e.Translate(trans)})
			}
		}
	}
	return validateResponse
}
