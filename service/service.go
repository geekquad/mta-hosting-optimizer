package service

import (
	"brevo/model"
	"fmt"
	"log"
	"os"
	"strconv"
)

type IPMetadata struct {
	IP    string
	Index int
}

type Result struct {
	Hostname string
}

type Service interface {
	GetActiveMTAsAboveThreshold() ([]Result, error)
	GetThresholdFromEnv() int
}

func GetActiveMTAsAboveThreshold() ([]Result, error) {
	threshold := GetThresholdFromEnv()
	log.Println("THRESHOLD", threshold)

	ipConfigs, err := model.GetServerInformation()
	if err != nil {
		log.Println("Unable to get server information")
	}

	activeIPs := make(map[string][]IPMetadata)
	for i, ipConfig := range ipConfigs {
		_, foundHostName := activeIPs[ipConfig.Hostname]
		if !foundHostName {
			activeIPs[ipConfig.Hostname] = []IPMetadata{}
		}

		if !ipConfig.Active {
			continue
		}

		metaData := IPMetadata{
			IP:    ipConfig.IP,
			Index: i,
		}

		activeIPs[ipConfig.Hostname] = append(activeIPs[ipConfig.Hostname], metaData)
	}

	var result []Result
	for hostname, activeIP := range activeIPs {
		if len(activeIP) <= threshold {
			result = append(result, Result{Hostname: hostname})
		}
	}

	return result, nil
}

func GetThresholdFromEnv() int {
	thresholdStr := os.Getenv("THRESHOLD")
	if thresholdStr == "" {
		return 1
	}
	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		fmt.Println("Error converting threshold to integer:", err)
		return 1
	}
	return threshold
}
