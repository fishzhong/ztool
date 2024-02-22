package trans

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

func init() {
	translator := zh.New()
	uni := ut.New(translator)
	trans, _ = uni.GetTranslator("zh")
}

func InitTrans() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = zh2.RegisterDefaultTranslations(v, trans)
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
			// skip if tag key says it should be ignored
			if name == "-" {
				return ""
			}
			return name
		})

	}
}

func Error(err error) string {
	for _, e := range err.(validator.ValidationErrors) {
		errM := e.Translate(trans)
		return errM
	}
	return ""
}

func Translate() {
	// 定义结构体
	type User struct {
		Name  string `validate:"required,min=2,max=20" label:"姓名"`
		Email string `validate:"required,email" label:"-"`
	}

	// 设置中文名称
	v := validator.New()
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	// 验证结构体
	user := User{Name: "", Email: "INVALID-EMAIL"}
	err := v.Struct(user)

	// 获取错误信息
	for _, er := range err.(validator.ValidationErrors) {
		fmt.Printf("%s 不合法\n", er.Field())
	}
}
