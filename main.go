package main

import (
	"github.com/paulofeitor/unique-ips/metrics"
	"github.com/paulofeitor/unique-ips/server"
)

func main() {
	met := metrics.New()

	met.Start(9102)

	s := server.New(met)

	s.Start(5000)
}
