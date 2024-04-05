package models

type Order struct {
	Id     int64 `json:"id" gorm:"primaryKey"`
	Price  int64 `json:"price"`
	MenuId int64 `json:"menu_id"`
	UserId int64 `json:"user_id"`
}
