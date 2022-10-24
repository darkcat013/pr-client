package domain

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"sync/atomic"

	"github.com/darkcat013/pr-client/config"
	"github.com/darkcat013/pr-client/utils"
	"go.uber.org/zap"
)

type Client struct {
	Id int
}

func NewClient() {
	atomic.AddInt64(&CliendId, 1)
	c := Client{
		Id: int(atomic.LoadInt64(&CliendId)),
	}
	utils.Log.Info("Client created", zap.Int("clientId", c.Id))

	go c.StartClient()
}

func (c *Client) StartClient() {
	utils.Log.Info("Client started", zap.Int("clientId", c.Id))

	utils.Log.Info("Get menu", zap.Int("clientId", c.Id))

	menu := c.getMenu()

	utils.Log.Info("Menu received", zap.Int("clientId", c.Id), zap.Any("menu", menu))

	utils.Log.Info("Start creating order", zap.Int("clientId", c.Id))

	foodOrderPayload := FoodOrderPayload{
		ClientId: c.Id,
		Orders:   make([]OrderPayload, 0),
	}

	for i := 0; i < menu.Restaurants; i++ {
		if rand.Intn(2) > 0 {
			order := c.newRestaurantOrder(menu.RestaurantsData[i].Menu, menu.RestaurantsData[i].RestaurantId)
			foodOrderPayload.Orders = append(foodOrderPayload.Orders, order)
		}
	}

	utils.Log.Info("Finished creating order", zap.Int("clientId", c.Id), zap.Any("order", foodOrderPayload))

	utils.Log.Info("Send order to food ordering service", zap.Int("clientId", c.Id))

	response := c.sendOrder(foodOrderPayload)

	utils.Log.Info("Received response from food ordering service", zap.Int("clientId", c.Id), zap.Any("orderResponse", response))

	// sort.Slice(response.Orders, func(i, j int) bool {
	// 	return response.Orders[i].EstimatedWaitingTime < response.Orders[j].EstimatedWaitingTime
	// })

	// for i := 0; i < len(response.Orders); i++ {
	// 	currentWaitTime := response.Orders[i].EstimatedWaitingTime
	// 	utils.Log.Info("Waiting for estimated time", zap.Int("clientId", c.Id), zap.Float64("estimated_wait_time", currentWaitTime), zap.Int("restaurantId", response.Orders[i].RestaurantId))
	// 	utils.SleepFor(currentWaitTime)

	// 	utils.Log.Info("Go pickup order", zap.Int("clientId", c.Id), zap.Int("restaurantId", response.Orders[i].RestaurantId))

	// 	utils.SleepBetween(config.CLIENT_ON_ROAD_TIME_MIN, config.CLIENT_ON_ROAD_TIME_MAX)
	// 	c.PickupOrder(response.Orders[i].RestaurantAddress)

	// 	utils.Log.Info("Order picked up", zap.Int("clientId", c.Id), zap.Int("restaurantId", response.Orders[i].RestaurantId))
	// }

	ClientDestroyedChan <- c.Id
}

func (c *Client) getMenu() Menu {
	resp, err := http.Get(config.FOOD_ORDERING_SERVICE_URL + "/menu")

	if err != nil {
		utils.Log.Fatal("Failed to get menu", zap.String("error", err.Error()), zap.Int("clientId", c.Id))
	} else {
		utils.Log.Info("Response from food ordering service | getMenu()", zap.Int("statusCode", resp.StatusCode), zap.Int("clientId", c.Id))
	}

	var menu Menu
	err = json.NewDecoder(resp.Body).Decode(&menu)

	if err != nil {
		utils.Log.Fatal("Failed to decode menu", zap.String("error", err.Error()), zap.Int("clientId", c.Id))
	}

	return menu
}

func (c *Client) newRestaurantOrder(menu []Food, restaurantId int) OrderPayload {

	foodsCount := rand.Intn(config.MAX_FOOD_PER_RESTAURANT) + 1
	var items []int
	var maxPrepTime float64

	for len(items) < foodsCount {

		randomFood := menu[rand.Intn(len(menu))]
		if maxPrepTime < randomFood.PreparationTime {
			maxPrepTime = randomFood.PreparationTime
		}

		items = append(items, randomFood.Id)
	}

	order := OrderPayload{
		RestaurantId: restaurantId,
		Items:        items,
		Priority:     rand.Intn(5) + 1,
		MaxWait:      maxPrepTime * config.MAX_PREP_TIME_COEFF,
		CreatedTime:  utils.GetCurrentTimeFloat(),
	}

	return order
}

func (c *Client) sendOrder(order FoodOrderPayload) FoodOrderResponse {
	body, err := json.Marshal(order)
	if err != nil {
		utils.Log.Fatal("Failed to convert order to JSON ", zap.String("error", err.Error()), zap.Any("order", order))
	}

	resp, err := http.Post(config.FOOD_ORDERING_SERVICE_URL+"/order", "application/json", bytes.NewBuffer(body))

	if err != nil {
		utils.Log.Fatal("Failed to send order to food order service", zap.String("error", err.Error()), zap.Any("order", order))
	} else {
		utils.Log.Info("Response from food order service", zap.Int("statusCode", resp.StatusCode), zap.Int("order", c.Id))
	}

	var result FoodOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		utils.Log.Fatal("Failed to decode FoodOrderResponse", zap.String("error", err.Error()), zap.Int("clientId", c.Id))
	}

	return result
}

// func (c *Client) PickupOrder(address string) DiningHallOrder {

// }
