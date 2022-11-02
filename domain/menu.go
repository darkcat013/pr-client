package domain

type Menu struct {
	Restaurants     int                    `json:"restaurants"`
	RestaurantsData map[int]RestaurantData `json:"restaurants_data"`
}

type RestaurantData struct {
	Name      string  `json:"name"`
	MenuItems int     `json:"menu_items"`
	Menu      []Food  `json:"menu"`
	Rating    float64 `json:"rating"`
}

type Food struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	PreparationTime  float64 `json:"preparation-time"`
	Complexity       int     `json:"complexity"`
	CookingApparatus string  `json:"cooking-apparatus"`
}
