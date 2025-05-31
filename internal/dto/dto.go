package dto

type CreateProductInput struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

type CreateUserInput struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type GetJWTInput struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"Name" validate:"required,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}
