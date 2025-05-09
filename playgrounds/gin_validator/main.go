// main.go
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := setupRouter()
	r.Run("0.0.0.0:8080")
}

type pathAndHandlerFuncPair struct {
	Path string
	gin.HandlerFunc
}

var handlerFuncs = []pathAndHandlerFuncPair{
	// c.BindJSON: 自動的にエラー時にレスポンスを返す（内部で c.Abort() を呼び出す）
	{"/users/bind_json_by_binding_tag_with_json_resp", newHandlerFunc(FuncType_BindJSON, TagType_Binding, ResponseType_JSON)},
	{"/users/bind_json_by_binding_tag_with_text_resp", newHandlerFunc(FuncType_BindJSON, TagType_Binding, ResponseType_Text)},
	{"/users/bind_json_by_binding_tag_without_resp", newHandlerFunc(FuncType_BindJSON, TagType_Binding, ResponseType_None)},
	{"/users/bind_json_by_validate_tag_with_json_resp", newHandlerFunc(FuncType_BindJSON, TagType_Validate, ResponseType_JSON)},
	{"/users/bind_json_by_validate_tag_with_text_resp", newHandlerFunc(FuncType_BindJSON, TagType_Validate, ResponseType_Text)},
	{"/users/bind_json_by_validate_tag_without_resp", newHandlerFunc(FuncType_BindJSON, TagType_Validate, ResponseType_None)},

	// c.ShouldBindJSON: エラーを返すだけなので自前でハンドリングできる
	{"/users/should_bind_json_by_binding_tag_with_json_resp", newHandlerFunc(FuncType_ShouldBindJSON, TagType_Binding, ResponseType_JSON)},
	{"/users/should_bind_json_by_binding_tag_with_text_resp", newHandlerFunc(FuncType_ShouldBindJSON, TagType_Binding, ResponseType_Text)},
	{"/users/should_bind_json_by_binding_tag_without_resp", newHandlerFunc(FuncType_ShouldBindJSON, TagType_Binding, ResponseType_None)},
	{"/users/should_bind_json_by_validate_tag_with_json_resp", newHandlerFunc(FuncType_ShouldBindJSON, TagType_Validate, ResponseType_JSON)},
	{"/users/should_bind_json_by_validate_tag_with_text_resp", newHandlerFunc(FuncType_ShouldBindJSON, TagType_Validate, ResponseType_Text)},
	{"/users/should_bind_json_by_validate_tag_without_resp", newHandlerFunc(FuncType_ShouldBindJSON, TagType_Validate, ResponseType_None)},
}

func setupRouter() *gin.Engine {
	// アクセスログ邪魔なので、アクセスログが出ないようにロガーを設定するgin.Default()ではなくて、gin.New()からRecoveryだけ設定
	r := gin.New()
	r.Use(gin.Recovery())

	// register HandlerFunc
	for _, pair := range handlerFuncs {
		r.POST(pair.Path, pair.HandlerFunc)
	}

	return r
}

type ResponseType int

const (
	ResponseType_None ResponseType = iota
	ResponseType_JSON
	ResponseType_Text
)

func handleBind(c *gin.Context, handleBindFunc func() error, responseType ResponseType) {

	if err := handleBindFunc(); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			// バリデーションエラー
			switch responseType {
			case ResponseType_JSON:
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // json
			case ResponseType_Text:
				c.String(http.StatusBadRequest, err.Error()) // text
			default:
				// no override
			}
			return
		}

		var gerr *gin.Error
		if errors.As(err, &gerr) {
			// バリデーション以外のginのカスタムエラー
			switch responseType {
			case ResponseType_JSON:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // json
			case ResponseType_Text:
				c.String(http.StatusInternalServerError, err.Error()) // text
			default:
				// no override
			}
			return
		}

		// 想定外の内部エラー
		switch responseType {
		case ResponseType_JSON:
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("unknown error: %s", err.Error())}) // json
		case ResponseType_Text:
			c.String(http.StatusInternalServerError, fmt.Sprintf("unknown error: %s", err.Error())) // text
		default:
			// no override
		}
		return
	}

	// 成功
	// 何も書かない=200なので、書き換わったことがわかりやすいよう200以外のものを選んだ → 201
	switch responseType {
	case ResponseType_JSON:
		c.JSON(http.StatusCreated, gin.H{"status": "ok"}) // json
	case ResponseType_Text:
		c.String(http.StatusCreated, "ok")
	default:
		// no override
		c.Status(http.StatusCreated)
	}
}

type FuncType int

const (
	FuncType_BindJSON = iota
	FuncType_ShouldBindJSON
)

type TagType int

const (
	TagType_Binding = iota
	TagType_Validate
)

type CreateUserRequestWithBindingTag struct {
	Name  string `json:"name,omitempty"  binding:"required"`
	Email string `json:"email,omitempty" binding:"required,email"`
}

type CreateUserRequestWithValidateTag struct {
	Name  string `json:"name,omitempty"  validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
}

func newHandlerFunc(funcType FuncType, tagType TagType, responseType ResponseType) gin.HandlerFunc {
	return func(c *gin.Context) {
		handleBind(
			c,
			func() error {
				switch funcType {
				// c.BindJSON
				case FuncType_BindJSON:
					switch tagType {
					case TagType_Binding:
						var req CreateUserRequestWithBindingTag
						return c.BindJSON(&req)
					case TagType_Validate:
						var req CreateUserRequestWithValidateTag
						return c.BindJSON(&req)
					default:
						panic("invalid FuncType")
					}
				// c.ShouldBindJSON
				case FuncType_ShouldBindJSON:
					switch tagType {
					case TagType_Binding:
						var req CreateUserRequestWithBindingTag
						return c.ShouldBindJSON(&req)
					case TagType_Validate:
						var req CreateUserRequestWithValidateTag
						return c.ShouldBindJSON(&req)
					default:
						panic("invalid FuncType")
					}
				default:
					panic("invalid FuncType")
				}
			},
			responseType,
		)
	}
}
