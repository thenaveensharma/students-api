package types

type Student struct {
	Id    int64
	Name  string `validate:"required"`
	Age   uint8  `validate:"gte=0,lte=130"`
	Email string `validate:"required,email"`
}
