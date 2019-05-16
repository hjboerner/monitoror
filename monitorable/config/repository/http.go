package repository

import (
	"net/http"

	"github.com/monitoror/monitoror/monitorable/config/models"
)

func (cr *configRepository) GetConfigFromUrl(url string) (config *models.Config, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	config, err = GetConfig(resp.Body)
	return
}
