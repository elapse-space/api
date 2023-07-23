package models

type Order struct {
	Id          int32   `json:"id" gorm:"primaryKey"`
	UserID      int     `json:"userID"`
	OrderNumber string  `json:"orderNumber"`
	Date        string  `json:"date"`
	Status      string  `json:"status"`
	Items       []int   `json:"items"` // Идентификаторы купленных товаров
	Total       float64 `json:"total"`
}
