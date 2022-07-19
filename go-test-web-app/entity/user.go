package entity

type User struct {
	UserId  int    `db:"user_id"` // IDです
	Name    string `db:"name"`    // 名前です
	Address string `db:"address"` // アドレスです
}
