package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	httpClient "github.com/gobench-io/gobench/clients/http"
	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/executor/scenario"
)

func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   5,
			Rate: 1000,
			Fu:   f,
		},
		{
			Nu:   5,
			Rate: 1000,
			Fu:   f,
		},
		{
			Nu:   5,
			Rate: 1000,
			Fu:   f,
		},
	}
}

type requestBody struct {
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	URL       string    `json:"url"`
}

func f(ctx context.Context, vui int) {
	client, err := httpClient.NewHttpClient(ctx, "logs-handler")
	if err != nil {
		log.Println("create new client fail: " + err.Error())
		return
	}

	url := "http://localhost:5000/logs"

	timeout := time.After(2 * time.Minute)

	for {
		select {
		case <-timeout:
			return
		default:
			request := requestBody{
				Timestamp: time.Now(),
				IP:        gofakeit.IPv4Address(),
				URL:       gofakeit.URL(),
			}

			body, _ := json.Marshal(request)
			go client.Post(ctx, url, body, map[string]string{
				"Content-Type": "application/json",
			})
			dis.SleepRatePoisson(10)
		}
	}
}
