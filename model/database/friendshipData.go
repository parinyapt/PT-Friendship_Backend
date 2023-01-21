package modelDatabase

import "time"

type FriendshipData struct {
	ID        uint      `gorm:"column:data_id"`
	UUID      string    `gorm:"column:data_uuid"`
	AccountID string    `gorm:"column:data_account_uuid"`
	Image     string    `gorm:"column:data_image"`
	From      string    `gorm:"column:data_from"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (FriendshipData) TableName() string {
	return "ptfs_friendship_data"
}
