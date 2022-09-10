package api

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nachonievag/ip_proxy_api/country"

	_ "github.com/nachonievag/ip_proxy_api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func routes(service *country.CountryHandler, port string) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))

	//INFO:Documentation endpoint
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))

	r.Get("/countries/CH/top_ten_isp", service.GetTopTenISP)
	r.Get("/countries/{countryCode}/ip/count", service.CountIPs)
	r.Get("/ip/{ip}", service.GetIP)

	return r
}
