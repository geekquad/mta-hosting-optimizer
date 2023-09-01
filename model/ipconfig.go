package model

import (
	"brevo/base"
	"log"
)

type IPConfig struct {
	IP       string
	Hostname string
	Active   bool
}

func GetServerInformation() ([]IPConfig, error) {
	keys, err := base.DataStore.Keys("*").Result()
	if err != nil {
		log.Println("Error fetching keys from Redis:", err)
		return nil, nil
	}

	var results []IPConfig
	for _, hostname := range keys {
		metaData, err := base.DataStore.HGetAll(hostname).Result()
		if err != nil {
			log.Println("err", err)
			continue
		}
		for ip, active := range metaData {
			result := IPConfig{
				Hostname: hostname,
				IP:       ip,
				Active:   active == "true",
			}
			results = append(results, result)
		}
	}
	return results, nil
}
