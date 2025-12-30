package target

//go:generate ../bin/unittestgen
type InputParams struct {
	Name  string `validate:"required,min=3,max=10"`
	Age   int    `validate:"gte=0,lte=120"`
	Email string `validate:"required,email"`
}
