package v1

import (
	. "gin-blog/internal/pkg/common"
	"gin-blog/internal/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhcn "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

var (
	Validate = validator.New()         // 实例化验证器
	ZhCn     = zh.New()                // 获取中文翻译器
	Cnc      = ut.New(ZhCn, ZhCn)      // 设置成中文翻译器
	Trans, _ = Cnc.GetTranslator("zh") // 获取翻译字典
)

// 格式化返回结果
func formatData(result interface{}, pager Pagination) map[string]interface{} {
	data := make(map[string]interface{})
	data["list"] = result
	data["pager"] = pager

	return data
}

// 校验输入参数
func ValidateParams(c *gin.Context, params interface{}) bool {
	_ = zhcn.RegisterDefaultTranslations(Validate, Trans)
	err := Validate.Struct(params)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errs []string
			for _, e := range errors {
				errs = append(errs, e.Translate(Trans))
			}
			core.Response(gin.H{"errors": errs}, core.Code(http.StatusUnprocessableEntity), core.Message("params value is invalid")).Json(c)
			return false
		}
	}
	return true
}
