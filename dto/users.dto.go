package dto;

type CreateUserRequest struct {
    Username string `json:"username" form:"username"`
    Email    string `json:"email" form:"email"`
    Phone    string `json:"phone" form:"phone"`
    Password string `json:"password" form:"password"`
}