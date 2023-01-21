package modelDatabase

import "time"

type Account struct {
	ID        uint      `gorm:"column:account_id"`
	UUID      string    `gorm:"column:account_uuid"`
	Username  string    `gorm:"column:account_username"`
	Name      string    `gorm:"column:account_name"`
	Image     string    `gorm:"column:account_image"`
	Token     string    `gorm:"column:account_token"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Account) TableName() string {
	return "ptfs_account"
}
