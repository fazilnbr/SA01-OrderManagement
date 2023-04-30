package domain

type RecOrder struct {
	ID           string    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Status       string    `json:"status"`
	Item         []Item `json:"items" gorm:"foreignKey:Item_id;references:ID"`
	Total        float64   `json:"total"`
	CurrencyUnit string    `json:"currencyUnit"`
}

type RecItem struct {
	ID          string  `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
