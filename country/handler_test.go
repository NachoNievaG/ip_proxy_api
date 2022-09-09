package country

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	gomock "github.com/golang/mock/gomock"
	"github.com/nachonievag/ip_proxy_api/internal/web"
	"github.com/stretchr/testify/assert"
)

const (
	getISPEndpoint     = "/countries/CH/top_ten_isp"
	getCountIPEndpoint = "/countries/%s/ip/count"
	getIPEndpoint      = "/ip/%s"
)

func TestHandler_GetTopTenISPSwiss(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := NewMockCountryGateway(mockCtrl)

	countryHandler := NewCountriesHTTPService(m)

	type TestCases struct {
		name           string
		expectedOutput []ISPCount
		expectedError  error
		status         int
	}

	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnTopTenISP",
			expectedOutput: sampleISP(),
		},
		{
			name:           "InputIsOK_ReturnEmptyTopTenISP",
			expectedOutput: []ISPCount{},
		},
		{
			name:           "InputIsOK_ReturnNilOutput",
			expectedOutput: nil,
		},
		{
			name:           "GatewayFails_ReturnError",
			expectedOutput: nil,
			expectedError:  errors.New("unexpected error"),
			status:         http.StatusInternalServerError,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			expect := m.EXPECT().TopTenISPByCountryCode(gomock.Any(), "CH")

			if tt.expectedError != nil {
				expect.Return(nil, tt.expectedError).Times(1)
			} else {
				expect.Return(tt.expectedOutput, nil).Times(1)
			}
			// Set http state
			req, err := http.NewRequest("GET", getISPEndpoint, nil)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(countryHandler.GetTopTenISP)
			handler.ServeHTTP(rr, req)

			if tt.expectedError != nil {
				var response web.ResponseError
				assert.Equal(t, tt.status, rr.Code)

				err = json.NewDecoder(rr.Body).Decode(&response)
				assert.Nil(t, err)
				assert.NotNil(t, response)

				assert.Equal(t, tt.expectedError.Error(), response.Result)
				return
			}

			assert.Equal(t, http.StatusOK, rr.Code)
			var response []ISPCount

			err = json.NewDecoder(rr.Body).Decode(&response)
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedOutput, response)
		})
	}
}

func TestHandler_CountIP(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := NewMockCountryGateway(mockCtrl)

	countryHandler := NewCountriesHTTPService(m)

	type TestCases struct {
		name           string
		countryCode    string
		expectedOutput IPCount
		expectedError  error
		status         int
	}
	sampleIPCount := IPCount{
		Country:  "AR",
		Quantity: 10,
	}
	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnCount",
			countryCode:    sampleIPCount.Country,
			expectedOutput: sampleIPCount,
		},
		{
			name:           "InputIsOK_ReturnZeroForCountry",
			countryCode:    sampleIPCount.Country,
			expectedOutput: IPCount{sampleIPCount.Country, 0},
		},
		{
			name:           "GatewayFails_ReturnError",
			countryCode:    "CH",
			expectedOutput: IPCount{},
			expectedError:  errors.New("unexpected error"),
			status:         http.StatusInternalServerError,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			expect := m.EXPECT().CountIPsByCountryCode(gomock.Any(), tt.countryCode)

			if tt.expectedError != nil {
				expect.Return(IPCount{}, tt.expectedError).Times(1)
			} else {
				expect.Return(tt.expectedOutput, nil).Times(1)
			}

			// Set http state

			req, err := http.NewRequest("GET", getISPEndpoint, nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("countryCode", tt.countryCode)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(countryHandler.CountIPs)
			handler.ServeHTTP(rr, req)

			if tt.expectedError != nil {
				var response web.ResponseError
				assert.Equal(t, tt.status, rr.Code)

				err = json.NewDecoder(rr.Body).Decode(&response)
				assert.Nil(t, err)
				assert.NotNil(t, response)

				assert.Equal(t, tt.expectedError.Error(), response.Result)
				return
			}

			assert.Equal(t, http.StatusOK, rr.Code)
			var response IPCount

			err = json.NewDecoder(rr.Body).Decode(&response)
			assert.Nil(t, err)
			assert.NotNil(t, response)
		})
	}
}

func TestHandler_GetIPInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := NewMockCountryGateway(mockCtrl)

	countryHandler := NewCountriesHTTPService(m)

	type TestCases struct {
		name           string
		ip             IPAddress
		expectedOutput IPModel
		expectedError  error
		status         int
	}

	stringIP := "192.111.1.11"
	testIP, err := NewIPAddressFromString(stringIP)
	assert.Nil(t, err)

	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnIPModel",
			ip:             testIP,
			expectedOutput: sampleIPModel(testIP, false),
		},
		{
			name:           "GatewayFails_ReturnInternalServerError",
			ip:             testIP,
			expectedOutput: IPModel{},
			expectedError:  errors.New("unexpected error"),
			status:         http.StatusInternalServerError,
		},
		{
			name:           "GatewayFails_ReturnBadRequest",
			ip:             testIP,
			expectedOutput: IPModel{},
			expectedError:  ErrInvalidIP,
			status:         http.StatusBadRequest,
		},
		{
			name:           "IPIsNotFound_ReturnNotFound",
			ip:             testIP,
			expectedOutput: IPModel{},
			expectedError:  ErrNotFound,
			status:         http.StatusNotFound,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			expect := m.EXPECT().GetIPInfo(gomock.Any(), tt.ip.IP.String())

			if tt.expectedError != nil {
				expect.Return(IPResponse{}, tt.expectedError).Times(1)
			} else {
				expect.Return(tt.expectedOutput.ToResponse(), nil).Times(1)
			}

			// Set http state

			req, err := http.NewRequest("GET", getISPEndpoint, nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("ip", tt.ip.IP.String())

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(countryHandler.GetIP)
			handler.ServeHTTP(rr, req)

			if tt.expectedError != nil {
				var response web.ResponseError
				assert.Equal(t, tt.status, rr.Code)

				err = json.NewDecoder(rr.Body).Decode(&response)
				assert.Nil(t, err)
				assert.NotNil(t, response)

				assert.Equal(t, tt.expectedError.Error(), response.Result)
				return
			}

			assert.Equal(t, http.StatusOK, rr.Code)
			var response IPResponse

			err = json.NewDecoder(rr.Body).Decode(&response)
			assert.Nil(t, err)
			assert.NotNil(t, response)
		})
	}
}
