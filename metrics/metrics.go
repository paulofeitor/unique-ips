package metrics

import (
	"log"
	"strconv"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	e                *echo.Echo
	uniqueIPsCounter prometheus.Counter
}

func New() *metrics {
	e := echo.New()
	e.HideBanner = true

	customRegistry := prometheus.NewRegistry()
	uniqueIPsCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "unique_ip_addresses_total",
			Help: "Counter of unique IP addresses that visited the company's website",
		},
	)
	if err := customRegistry.Register(uniqueIPsCounter); err != nil {
		log.Fatal("failed to register counter:", err)
	}

	e.GET("/metrics", echoprometheus.NewHandlerWithConfig(echoprometheus.HandlerConfig{Gatherer: customRegistry}))

	return &metrics{
		e:                e,
		uniqueIPsCounter: uniqueIPsCounter,
	}
}

func (m *metrics) Start(port int) {
	go func() {
		m.e.Logger.Fatal(m.e.Start(":" + strconv.Itoa(port)))
	}()
}

func (m *metrics) UniqueIPsInc() {
	m.uniqueIPsCounter.Inc()
}
