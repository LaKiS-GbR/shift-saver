package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Admin    bool   `gorm:"not null"`
}

type Shift struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"not null"`
	StartTime string `gorm:"not null"`
	EndTime   string `gorm:"not null"`
	TotalTime string `gorm:"not null"`
	TotalPay  string `gorm:"not null"`
}
