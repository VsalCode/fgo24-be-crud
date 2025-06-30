package dto;

type CreateUserRequest struct {
    Username string `json:"username" form:"username"`
    Email    string `json:"email" form:"email"`
    Phone    string `json:"phone" form:"phone"`
    Password string `json:"password" form:"password"`
}

type UpdateUserRequest struct {
    Username *string `json:"username,omitempty"`
    Email    *string `json:"email,omitempty"`
    Phone    *string `json:"phone,omitempty"`
    Password *string `json:"password,omitempty"`
}
