package write

type OrderItem struct {
	Uuid        string `gorm:"index;unique"`
	OrderId     string
	ProductId   string
	Title       string
	Description string
	Price       int
}
