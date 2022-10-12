package domain

type FoodOrderPayload struct {
	ClientId int            `json:"client_id"`
	Orders   []OrderPayload `json:"orders"`
}

type OrderPayload struct {
	RestaurantId int     `json:"restaurant_id"`
	Items        []int   `json:"items"`
	Priority     int     `json:"priority"`
	MaxWait      float64 `json:"max_wait"`
	CreatedTime  float64 `json:"created_time"`
}

type FoodOrderResponse struct {
	OrderId int             `json:"order_id"`
	Orders  []OrderResponse `json:"orders"`
}

type OrderResponse struct {
	RestaurantId         int     `json:"restaurant_id"`
	RestaurantAddress    string  `json:"restaurant_address"`
	OrderId              int     `json:"order_id"`
	EstimatedWaitingTime float64 `json:"estimated_waiting_time"`
	CreatedTime          float64 `json:"created_time"`
	RegisteredTime       float64 `json:"registered_time"`
}
