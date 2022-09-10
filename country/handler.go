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

// GetTopTenISP  godoc
// @Summary      Retrieve Top 10 internet providers from Switzerland
// @Description  using the country code, obtain the top ten internet service providers from Switzerland, descending order.
// @Tags         Country
// @Accept       json
// @Produce      json
// @Success      200  {object}  []ISPCount
// @Failure      500  {object}  web.ResponseError
// @Router       /countries/CH/top_ten_isp [get]
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

// Count IP Addresses godoc
// @Summary      Given a country code, retrieve the IP count
// @Description  using the country code, count every ip given in the database.
// @Param        countryCode   path      string  true  "Country Code"
// @Tags         Country
// @Accept       json
// @Produce      json
// @Success      200  {object}  IPCount
// @Failure      500  {object}  web.ResponseError
// @Router       /countries/{countryCode}/ip/count [get]
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

// Get IP Info   godoc
// @Summary      Giving an IP as an input, return every detail related to it
// @Description  After receiving the IP, it is validated to ensure it is a proper address, then its content is used to retrieve information related to it, such as its country code.
// @Tags         Country
// @Accept       json
// @Produce      json
// @Param        ip   path      string  true  "IP address"
// @Success      200  {object}  IPResponse
// @Failure      400  {object}  web.ResponseError
// @Failure      404  {object}  web.ResponseError
// @Failure      500  {object}  web.ResponseError
// @Router       /ip/{ip} [get]
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
