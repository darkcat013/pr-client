package main

import (
	"github.com/darkcat013/pr-client/config"
	"github.com/darkcat013/pr-client/domain"
	"github.com/darkcat013/pr-client/utils"
	"go.uber.org/zap"
)

func main() {
	utils.InitializeLogger()

	for i := 0; i < config.MAX_CLIENTS; i++ {
		domain.NewClient()
		utils.SleepBetween(config.CLIENT_BETWEEN_TIME_MIN, config.CLIENT_BETWEEN_TIME_MAX)
	}

	for {
		destroyedClientId := <-domain.ClientDestroyedChan
		utils.SleepBetween(config.CLIENT_BETWEEN_TIME_MIN, config.CLIENT_BETWEEN_TIME_MAX)
		utils.Log.Info("Client destroyed", zap.Int("clientId", destroyedClientId))

		domain.NewClient()
	}
}
