package entity

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	IsActive *bool  `json:"is_active"`
	Role     Role   `json:"role"`
	Password string
}
