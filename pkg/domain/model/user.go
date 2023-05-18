package model

type User struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	DisplayName string `gorm:"type:varchar(10)" json:"display_name"`
	Email       string `gorm:"type:varchar(30);uniqueindex;not null" json:"email"`
	Password    string `gorm:"type:varchar(10);not null" json:"password"`
}
