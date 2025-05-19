package entity

type User struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	Email     string `gorm:"column:email;type:varchar(255);not null;unique"`
	Password  string `gorm:"column:password;type:varchar(255);not null"`
	CreatedAt string `gorm:"column:created_at;type:timestamp;not null;default:now()"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
}

func (e *User) TableName() string {
	return "user"
}
