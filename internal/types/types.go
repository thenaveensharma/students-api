package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" validate:"required "`
	Age   uint8  `json:"age" validate:"gte=0,lte=130"`
	Email string `json:"email" validate:"required,email"`
}
