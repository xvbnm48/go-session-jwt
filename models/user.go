package models

type User struct {
	Id          int64  `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	NamaLengkap string `grom:"varchar(300)" json:"nama_lengkap"`
	Username    string `gorm:"varchar(300)" json:"username"`
	Password    string `gorm:"varchar(300)" json:"password"`
}
