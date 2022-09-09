package country

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nachonievag/ip_proxy_api/internal/web"
)

type CountryHandler struct {
	gateway CountryGateway
}

func NewCountriesHTTPService(gtw CountryGateway) *CountryHandler {
	return &CountryHandler{
		gateway: gtw,
	}
}

func (h *CountryHandler) GetTopTenISP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countryCode := "CH"

	topTen, err := h.gateway.TopTenISPByCountryCode(ctx, countryCode)
	if err != nil {
		web.Failure(err, http.StatusInternalServerError).Send(w)
		return
	}

	web.Success(topTen, http.StatusOK).Send(w)
}

func (h *CountryHandler) CountIPs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countryCode := chi.URLParam(r, "countryCode")

	//TODO: Check if param is a proper country code
	count, err := h.gateway.CountIPsByCountryCode(ctx, countryCode)
	if err != nil {
		web.Failure(err, http.StatusInternalServerError).Send(w)
		return
	}

	web.Success(count, http.StatusOK).Send(w)
}

func (h *CountryHandler) GetIP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ip := chi.URLParam(r, "ip")

	info, err := h.gateway.GetIPInfo(ctx, ip)
	if err != nil {
		web.Failure(handleError(err)).Send(w)
		return
	}

	web.Success(info, http.StatusOK).Send(w)
}
