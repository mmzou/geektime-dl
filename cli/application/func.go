package application

import (
	"github.com/mmzou/geektime-dl/config"
	"github.com/mmzou/geektime-dl/service"
)

func getService() *service.Service {
	return config.Instance.ActiveUserService()
}
