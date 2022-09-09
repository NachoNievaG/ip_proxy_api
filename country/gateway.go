package country

import (
	"context"
	"fmt"
)

//go:generate mockgen -destination=mock_country_gateway.go -package=country -source=gateway.go CountryGateway

type CountryGateway interface {
	TopTenISPByCountryCode(ctx context.Context, countryCode string) ([]ISPCount, error)
	CountIPsByCountryCode(ctx context.Context, countryCode string) (IPCount, error)
	GetIPInfo(ctx context.Context, ip string) (IPResponse, error)
}

type countryGateway struct {
	repository CountryRepository
}

func NewCountryGateway(repo CountryRepository) CountryGateway {
	return &countryGateway{
		repository: repo,
	}
}

func (gtw *countryGateway) TopTenISPByCountryCode(ctx context.Context, countryCode string) ([]ISPCount, error) {
	return gtw.repository.TopISPByCountryCode(ctx, countryCode)
}

func (gtw *countryGateway) CountIPsByCountryCode(ctx context.Context, countryCode string) (IPCount, error) {
	return gtw.repository.CountIPs(ctx, countryCode)
}

func (gtw *countryGateway) GetIPInfo(ctx context.Context, ip string) (IPResponse, error) {
	ipAddr, err := NewIPAddressFromString(ip)
	if err != nil {
		return IPResponse{}, fmt.Errorf("NewIPAddress: %w", err)
	}

	ipInfo, err := gtw.repository.GetIP(ctx, ipAddr.IntIP)
	if err != nil {
		return IPResponse{}, fmt.Errorf("gtw.repository.GetIP: %w", err)
	}

	return ipInfo.ToResponse(), nil
}
