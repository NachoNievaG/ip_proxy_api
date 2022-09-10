package api

import (
	"github.com/nachonievag/ip_proxy_api/country"
)

func Start(port string) {
	db := ConnectToDB()
	defer db.Close()
	countryRepo := country.NewCountryRepository(db)
	countryGTW := country.NewCountryGateway(countryRepo)
	handler := country.NewCountriesHTTPService(countryGTW)

	r := routes(handler, port)
	server := newServer(port, r)

	server.Start()
}
