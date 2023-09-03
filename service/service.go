package service

import (
	"fmt"
	"log"
	"mta-hosting-optimizer/model"
	"os"
	"strconv"
)

type Result struct {
	Hostname string `json:"hostname"`
}

func GetActiveMTAsAboveThreshold() ([]Result, error) {
	threshold := GetThresholdFromEnv()

	ipConfigs, err := model.GetServerInformation()
	if err != nil {
		log.Println("Unable to get server information")
	}

	activeIPs := make(map[string][]string)
	for _, ipConfig := range ipConfigs {
		_, foundHostName := activeIPs[ipConfig.Hostname]
		if !foundHostName {
			activeIPs[ipConfig.Hostname] = []string{}
		}

		if !ipConfig.Active {
			continue
		}

		activeIPs[ipConfig.Hostname] = append(activeIPs[ipConfig.Hostname], ipConfig.IP)
	}

	result := []Result{}
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
