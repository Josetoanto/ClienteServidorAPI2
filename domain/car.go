package domain

type Car struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Brand  string  `json:"brand"`
	Model  string  `json:"model"`
	Year   int     `json:"year"`
	Precio float64 `json:"precio"` // Nuevo campo
}
