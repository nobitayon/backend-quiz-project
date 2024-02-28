package models

type User struct {
	Id       int    `gorm:"column:id" json:"id"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Role     string `gorm:"column:role" json:"-"`
}

func (User) TableName() string {
	return "user"
}
