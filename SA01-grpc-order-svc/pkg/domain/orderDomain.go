package domain

type Order struct {
	ID           string  `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Status       string  `json:"status"`
	Item_id      int     `json:"items"`
	Item         *Item   `json:"-" gorm:"foreignKey:Item_id;references:ID"`
	Total        float64 `json:"total"`
	CurrencyUnit string  `json:"currencyUnit"`
}

type Item struct {
	ID          string  `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
