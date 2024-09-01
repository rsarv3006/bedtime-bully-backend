package alert

import (
	"bedtime-bully-backend/config"

	"github.com/apialerts/apialerts-go"
)

func Connect() *apialerts.APIAlertsClient {
	apiKey := config.Config("API_ALERTS_KEY")
	client := apialerts.ApiAlertsClient(apiKey)

	return client
}
