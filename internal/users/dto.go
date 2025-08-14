package users

type CreateUserDTO struct {
	Username    string `json:"user_name" validate:"required,min=5,max=50"`
	DisplayName string `json:"display_name" validate:"required,min=5,max=50"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=5,max=50"`
	Color       string `json:"color" validate:"omitempty,hexcolor"`
}

type UpdateUserDTO struct {
	Username    string `json:"user_name" validate:"omitempty,min=5,max=50"`
	DisplayName string `json:"display_name" validate:"omitempty,min=5,max=50"`
	Email       string `json:"email" validate:"omitempty,email"`
	Password    string `json:"password" validate:"omitempty,min=5,max=50"`
	Color       string `json:"color" validate:"omitempty,hexcolor"`
}
