package entity

type Todo struct {
	ID            int    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	Todo          string `gorm:"column:todo;type:varchar(255)"`
	UpdatedAt     string `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
	CreatedAt     string `gorm:"column:created_at;type:timestamp;not null;default:now()"`
}

func (e *Todo) TableName() string {
	return "todo"
}
