package server

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Metrics interface {
	UniqueIPsInc()
}

type server struct {
	e      *echo.Echo
	unique *sync.Map
	m      Metrics
}

func New(met Metrics) *server {
	e := echo.New()
	e.HideBanner = true

	return &server{
		e:      e,
		unique: &sync.Map{},
		m:      met,
	}
}

func (s *server) Start(port int) error {
	s.e.POST("/logs", s.LogsHandler)

	return s.e.Start(":" + strconv.Itoa(port))
}

type logsRequest struct {
	IP string `json:"ip"`
}

func (s *server) LogsHandler(c echo.Context) error {
	request := logsRequest{}
	if err := c.Bind(&request); err != nil {
		log.Errorf("error binding request: %w", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	if _, loaded := s.unique.LoadOrStore(request.IP, true); !loaded {
		// not loaded means this IP was never seen by the service
		s.m.UniqueIPsInc()
	}

	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}
