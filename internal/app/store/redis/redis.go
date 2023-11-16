package redis

import (
	"fmt"
	"urbathon-backend-2023/pkg/config"
)

type Redis struct {
	Address string
	Port    string
}

func New() *Redis {
	return &Redis{
		Address: fmt.Sprintf("%s", config.GetConfig().Get("redis.address")),
		Port:    fmt.Sprintf("%s", config.GetConfig().Get("redis.port")),
	}
}
