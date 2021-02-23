package model

// UserModel UserModel
type UserModel struct {
	ID   int    `gorm:"id" json:"id"`
	UID  int32  `gorm:"uid" json:"uid"`
	Name string `gorm:"name" json:"name"`
	Sex  string `gorm:"sex" json:"sex"`
}
