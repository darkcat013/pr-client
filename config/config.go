package config

import "time"

const LOGS_ENABLED = true

const PORT = "8089"

// const FOOD_ORDERING_SERVICE_URL = "http://localhost:8088"
const FOOD_ORDERING_SERVICE_URL = "http://host.docker.internal:8088"

const TIME_UNIT = time.Millisecond * TIME_UNIT_COEFF
const TIME_UNIT_COEFF = 100

const MAX_CLIENTS = 3
const CLIENT_BETWEEN_TIME_MIN = 2
const CLIENT_BETWEEN_TIME_MAX = 4
const CLIENT_ORDERING_TIME_MIN = 2
const CLIENT_ORDERING_TIME_MAX = 4
const CLIENT_ON_ROAD_TIME_MIN = 2
const CLIENT_ON_ROAD_TIME_MAX = 4

const MAX_FOOD_PER_RESTAURANT = 5
const MAX_PREP_TIME_COEFF = 1.3
