package ims

type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	MobilePhone string `json:"mobilePhone"`
	ModifiedAt  string `json:"modifiedAt"`
	UserID      int    `json:"userID"`
	CreatedAt   string `json:"createdAt"`
	Email       string `json:"email"`
}

type ChangePasswordRequest struct {
	ExistingPassword string `json:"existingPassword"`
	NewPassword      string `json:"newPassword"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type SignupRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}
