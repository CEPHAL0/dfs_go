package enums

type OrderStatus string

const (
	ORDERED    OrderStatus = "Ordered"
	PROCESSING OrderStatus = "Processing"
	DELIVERED  OrderStatus = "Delivered"
	MISSED     OrderStatus = "Missed"
)
