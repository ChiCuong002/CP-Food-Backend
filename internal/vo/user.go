package vo

type UserRegistrationRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UserLoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


// response
type UserResponse struct {
	ID         int32          `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Status     interface{} `json:"status"`
}

type UserRegisterResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Tokens TokensResponse `json:"tokens"`
}
type TokensResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type UserLoginResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Tokens TokensResponse `json:"tokens"`
}
type ListUsersResponse struct {
	Limit int `json:"limit"`
	Page int `json:"page"`
	TotalRows int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
	Rows interface{} `json:"rows"`
}