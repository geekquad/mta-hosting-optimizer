package main

import (
	"fmt"
	"mta-hosting-optimizer/base"
	"net/http"
	"testing"
	"time"
)

const url = "http://localhost:8080/hostname"

func BenchmarkRedisSet(b *testing.B) {
	base.InitRedis()
	go startServer()
	time.Sleep(2 * time.Second)

	for i := 0; i < b.N; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making GET request: %v\n", err)
			return
		}
		defer resp.Body.Close()
	}
}

func BenchmarkKeyDBSet(b *testing.B) {
	base.InitKeyDB()
	go startServer()
	time.Sleep(2 * time.Second)

	for i := 0; i < b.N; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making GET request: %v\n", err)
			return
		}
		defer resp.Body.Close()
	}
}
