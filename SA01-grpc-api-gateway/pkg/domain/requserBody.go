package domain

type Order struct {
	ID           string  `json:"id" `
	Status       string  `json:"status"`
	Item         []Item  `json:"items"`
	Total        float64 `json:"total"`
	CurrencyUnit string  `json:"currencyUnit"`
}

type Item struct {
	ID          string  `json:"id" `
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type UpdateOrder struct {
	ID     string `json:"id" `
	Status string `json:"status"`
}

type Filter struct {
	Status    string  `json:"status"`
	MinTotal  float64 `json:"mintotal"`
	MaxTotal  float64 `json:"maxtotal"`
	SortBy    string  `json:"sortby"`
	SortOrder string  `json:"sortorder"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
