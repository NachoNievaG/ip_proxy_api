package country

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

//go:generate mockgen -destination=mock_country_repo.go -package=country -source=repository.go CountryRepository

type CountryRepository interface {
	TopISPByCountryCode(ctx context.Context, countryCode string) ([]ISPCount, error)
	CountIPs(ctx context.Context, countryCode string) (IPCount, error)
	GetIP(ctx context.Context, ip uint32) (IPModel, error)
}

type countryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) CountryRepository {
	return &countryRepository{
		db: db,
	}
}

// INFO: countIPsQuery shall return the ∑ of ips given a certain countryCode. Feels automagical. Hmm...
const (
	topTenISPQuery = "SELECT isp ,count(1) as count FROM ip2location_px7 where country_code='CH' group by isp order by count DESC limit 10;"
	countIPsQuery  = "SELECT country_code, SUM(greatest( ip_to-ip_from, 1 )) as sum from ip2location_px7 where country_code= $1 group by country_code;"
	getIPInfoQuery = "SELECT ip_from, ip_to, proxy_type, country_code, country_name, region_name, city_name, isp, domain, usage_type, asn, asn FROM ip2location_px7 where $1 >= ip_from and ip_to >= $1;"
)

// INFO: Sentinel Errors
var (
	ErrNotFound = errors.New("required information was not found")
)

func (repo *countryRepository) TopISPByCountryCode(ctx context.Context, countryCode string) ([]ISPCount, error) {
	var result []ISPCount
	out, err := repo.db.Query(topTenISPQuery)
	if err != nil {
		return nil, err
	}

	for out.Next() {
		var ispCount ISPCount
		if err := out.Scan(&ispCount.ISP, &ispCount.Quantity); err != nil {
			return nil, err
		}

		result = append(result, ispCount)
	}

	err = out.Close()
	if err != nil {
		return nil, err
	}

	if err := out.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *countryRepository) CountIPs(ctx context.Context, countryCode string) (IPCount, error) {
	var result IPCount

	out := repo.db.QueryRowContext(ctx, countIPsQuery, countryCode)

	if err := out.Scan(&result.Country, &result.Quantity); err != nil {
		//INFO: Avoid No Rows error so bad msg is mitigated
		if errors.Is(err, sql.ErrNoRows) {
			return IPCount{
				Country:  countryCode,
				Quantity: 0,
			}, nil
		}

		return IPCount{}, err
	}

	if err := out.Err(); err != nil {
		return IPCount{}, err
	}

	return result, nil
}

func (repo *countryRepository) GetIP(ctx context.Context, ip uint32) (IPModel, error) {
	var result IPModel
	out := repo.db.QueryRowContext(ctx, getIPInfoQuery, ip)

	if err := out.Scan(
		&result.IPFrom,
		&result.IPTo,
		&result.ProxyType,
		&result.CountryCode,
		&result.CountryName,
		&result.RegionName,
		&result.CityName,
		&result.ISP,
		&result.Domain,
		&result.UsageType,
		&result.ASN,
		&result.AS,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return IPModel{}, ErrNotFound
		}

		return IPModel{}, fmt.Errorf("GetIP Scan: %w", err)
	}

	if err := out.Err(); err != nil {
		return IPModel{}, fmt.Errorf("GetIP row: %w", err)
	}

	return result, nil
}
