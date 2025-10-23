package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand/v2"
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
	}
}

var ips = []string{
	"83.150.59.250",
	"83.150.59.251",
	"83.150.59.252",
	"83.150.59.253",
	"83.150.59.254",
	"83.150.59.255",
	"83.150.59.256",
	"83.150.59.257",
	"83.150.59.258",
	"83.150.59.259",
	"83.150.59.260",
}

type requestBody struct {
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	URL       string    `json:"url"`
	Field1    string    `json:"field1"`
	Field2    string    `json:"field2"`
	Field3    string    `json:"field3"`
	Field4    string    `json:"field4"`
	Field5    string    `json:"field5"`
	Field6    string    `json:"field6"`
	Field7    string    `json:"field7"`
	Field8    string    `json:"field8"`
	Field9    string    `json:"field9"`
	Field10   string    `json:"field10"`
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
				IP:        ips[rand.IntN(10)],
				URL:       gofakeit.URL(),
				Field1:    gofakeit.Phrase(),
				Field2:    gofakeit.Phrase(),
				Field3:    gofakeit.Phrase(),
				Field4:    gofakeit.Phrase(),
				Field5:    gofakeit.Phrase(),
				Field6:    gofakeit.Phrase(),
				Field7:    gofakeit.Phrase(),
				Field8:    gofakeit.Phrase(),
				Field9:    gofakeit.Phrase(),
				Field10:   gofakeit.Phrase(),
			}

			body, _ := json.Marshal(request)
			go client.Post(ctx, url, body, map[string]string{
				"Content-Type": "application/json",
			})
			dis.SleepRatePoisson(10)
		}
	}
}
