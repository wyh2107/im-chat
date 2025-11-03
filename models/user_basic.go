package models

type UserBasic struct {
}

func (u *UserBasic) TableName() string {
	return "user_basic"
}
