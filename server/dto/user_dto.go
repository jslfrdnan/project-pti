package dto

type UserRequest struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserData struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserResponse struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       UserData `json:"data,omitempty"`
}