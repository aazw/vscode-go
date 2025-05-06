package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/aazw/vscode-go/playgrounds/custom_error_with_stacktrace/cerrors"
	"github.com/aazw/vscode-go/playgrounds/go-playground_validator_with_cerrors/validatorx"
)

// User contains user information
type User struct {
	FirstName      string     `json:"first_name"     validate:"required"`
	LastName       string     `json:"last_name"      validate:"required"`
	Age            uint8      `json:"age"            validate:"gte=0,lte=130"`
	Email          string     `json:"email"          validate:"required,email"`
	Gender         string     `json:"gender"         validate:"oneof=male female prefer_not_to"`
	FavouriteColor string     `json:"favorite_color" validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `json:"addresses"      validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city"   validate:"required"`
	Planet string `json:"planet" validate:"required"`
	Phone  string `json:"phone"  validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New(validator.WithRequiredStructEnabled())

	// JSON タグ名を優先的に返す関数を登録
	// → namespace、fieldがjsonタグ型の名前になる
	//    これを登録しないと、namespace、fieldにはGoのフィールド名が使われる = namespaceとstruct_namespaceが一緒、fieldとstruct_fieldが一緒になる
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// `json:"-"` のときは空文字を返し、Field() では飛ばす
		tag := fld.Tag.Get("json")
		if tag == "-" || tag == "" {
			return fld.Name // 代替として Go フィールド名
		}
		// オプション（`,omitempty` など）を除去して最初のトークンだけ
		name := strings.Split(tag, ",")[0]
		return name
	})

	validateStruct()

	fmt.Printf("\n")

	validateVariable()
}

func validateStruct() {

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Gender:         "male",
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Printf("error: %+v\n", err)
			return
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				fmt.Printf("namespace: %v\n", e.Namespace())
				fmt.Printf("field: %v\n", e.Field())
				fmt.Printf("struct_namespace: %v\n", e.StructNamespace())
				fmt.Printf("struct_field: %v\n", e.StructField())
				fmt.Printf("tag: %v\n", e.Tag())
				fmt.Printf("actual_tag: %v\n", e.ActualTag())
				fmt.Printf("kind: %v\n", e.Kind())
				fmt.Printf("type: %v\n", e.Type())
				fmt.Printf("value: %v\n", e.Value())
				fmt.Printf("param: %v\n", e.Param())
				fmt.Printf("\n")
			}
		}

		// cerrors
		options := []cerrors.Option{
			cerrors.WithCause(err),
		}
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {

				msg := validatorx.Message(reflect.ValueOf(user), e)
				fmt.Printf("msg: %s\n", msg)

				options = append(options, cerrors.WithMessage(msg))
			}
		}
		cerr := cerrors.ErrUnknown.New(options...)

		// slog
		fmt.Printf("\n")
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
		logger := slog.New(handler)
		logger.Error("error message", "err", cerr)
	}

	// save user to database
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
