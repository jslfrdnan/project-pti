package dto

type TodoRequest struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Todo          string `json:"todo"`
	UpdatedAt     string `json:"updated_at"`
	CreatedAt     string `json:"created_at"`
}

type TodoData struct {
	ID            int    `json:"id"`
	Todo          string `json:"todo"`
}

type TodoResponse struct {
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Data       TodoData  `json:"data,omitempty"`
}

// Error implements error.
func (i *TodoResponse) Error() string {
	panic("unimplemented")
}
