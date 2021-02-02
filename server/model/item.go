package model

// User struct
type User struct {
	ID           uint64 `json:"id"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	RegisterDate string `json:"register_date"`
}

// Item struct
type Item struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

// Items struct
type Items struct {
	Items []Item `json:"items"`
}
