package dto

type CreateUserRequest struct {
	FullName        string `json:"fullname"`
	Username        string `json:"username"`
	UserStatus      string `json:"user_status"`
	UserRole        string `json:"user_role"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UpdateUserRequest struct {
	FullName   string `json:"fullname"`
	Username   string `json:"username"`
	UserStatus string `json:"user_status"`
	UserRole   string `json:"user_role"`
}

type ChangeUserPassword struct {
	NewPassword string `json:"new_password"`
}
