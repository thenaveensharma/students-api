package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   uint8  `json:"age" validate:"required,gte=0,lte=130"`
}
