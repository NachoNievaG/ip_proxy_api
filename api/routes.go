package api

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nachonievag/ip_proxy_api/country"
)

func routes(service *country.CountryHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))

	// define routes
	r.Get("/countries/CH/top_ten_isp", service.GetTopTenISP)
	r.Get("/countries/{countryCode}/ip/count", service.CountIPs)
	r.Get("/ip/{ip}", service.GetIP)

	return r
}
