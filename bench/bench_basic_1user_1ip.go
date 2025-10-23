package main

import (
	"context"
	"log"
	"time"

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
	}
}

var jsonInput = `
{
    "timestamp": "2020-06-24T15:27:00.123456Z",
    "ip": "83.150.59.250",
    "url": "google.com"
}
`

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
			go client.Post(ctx, url, []byte(jsonInput), map[string]string{
				"Content-Type": "application/json",
			})
			dis.SleepRatePoisson(10)
		}
	}
}
