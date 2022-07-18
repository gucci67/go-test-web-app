package entity

type User struct {
	UserId  int    `gorm:"primaryKey;AUTO_INCREMENT"` // IDです
	Name    string `gorm:"not null"`                  // 名前です
	Address string `gorm:"not null"`                  // アドレスです
}
