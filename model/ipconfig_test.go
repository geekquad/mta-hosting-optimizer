package model

import (
	"mta-hosting-optimizer/base"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerInformation(t *testing.T) {
	base.InitRedis()

	err := base.DataStore.HMSet("mta-prod-1", map[string]interface{}{"127.0.0.1": "true", "127.0.0.2": "false"}).Err()
	if err != nil {
		t.Fatal(err)
	}
	err = base.DataStore.HMSet("mta-prod-2", map[string]interface{}{"127.0.0.3": "true", "127.0.0.4": "true", "127.0.0.5": "false"}).Err()
	if err != nil {
		t.Fatal(err)
	}
	err = base.DataStore.HMSet("mta-prod-3", map[string]interface{}{"127.0.0.6": "false"}).Err()
	if err != nil {
		t.Fatal(err)
	}

	results, err := GetServerInformation()

	assert.NoError(t, err)
	expectedResults := []IPConfig{
		{Hostname: "mta-prod-1", IP: "127.0.0.1", Active: true},
		{Hostname: "mta-prod-1", IP: "127.0.0.2", Active: false},
		{Hostname: "mta-prod-2", IP: "127.0.0.3", Active: true},
		{Hostname: "mta-prod-2", IP: "127.0.0.4", Active: true},
		{Hostname: "mta-prod-2", IP: "127.0.0.5", Active: false},
		{Hostname: "mta-prod-3", IP: "127.0.0.6", Active: false},
	}

	sort.Slice(expectedResults, func(i, j int) bool {
		return expectedResults[i].Hostname < expectedResults[j].Hostname || (expectedResults[i].Hostname == expectedResults[j].Hostname && expectedResults[i].IP < expectedResults[j].IP)
	})
	sort.Slice(results, func(i, j int) bool {
		return results[i].Hostname < results[j].Hostname || (results[i].Hostname == results[j].Hostname && results[i].IP < results[j].IP)
	})

	// Compare the actual results with the expected results
	assert.Equal(t, expectedResults, results)
}
