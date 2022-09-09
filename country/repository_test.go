package country

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepository_TopTenISP(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	repo := NewCountryRepository(db)
	columns := []string{"isp", "count"}
	type TestCases struct {
		name           string
		countryCode    string
		expectedOutput []ISPCount
		expectedError  error
		scanError      bool
		closeError     bool
		rowError       bool
	}

	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnTopTenISP",
			countryCode:    "CH",
			expectedOutput: sampleISP(),
			expectedError:  nil,
		},
		{
			name:           "ConnectionError_ReturnError",
			countryCode:    "CH",
			expectedOutput: nil,
			expectedError:  errors.New("unexpected error"),
		},
		{
			name:           "ScanError_ReturnError",
			countryCode:    "CH",
			expectedOutput: nil,
			expectedError:  nil,
			scanError:      true,
		},
		{
			name:           "CloseError_ReturnError",
			countryCode:    "CH",
			expectedOutput: nil,
			expectedError:  nil,
			closeError:     true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			rows := sqlmock.NewRows(columns)
			for _, ispCount := range tt.expectedOutput {
				rows.AddRow(ispCount.ISP, ispCount.Quantity)
			}

			//NOTE:Using QuoteMeta() given that ExpectQuery expects a regexp
			exec := mock.ExpectQuery(regexp.QuoteMeta(topTenISPQuery)).WithArgs()

			if tt.closeError {
				rows.CloseError(sql.ErrConnDone)
			}

			if tt.scanError {
				rows.AddRow("example", "bad value")
			}

			if tt.expectedError != nil {
				exec.WillReturnError(tt.expectedError)
			}

			exec.WillReturnRows(rows)

			res, err := repo.TopISPByCountryCode(context.Background(), tt.countryCode)
			if tt.expectedError != nil || tt.rowError || tt.scanError || tt.closeError {
				assert.Nil(t, res)
				assert.Error(t, err)
				return
			}
			assert.NotNil(t, res)
			assert.Equal(t, len(tt.expectedOutput), len(res))
			for i := range res {
				assert.Equal(t, tt.expectedOutput[i], res[i])
			}
			assert.Nil(t, mock.ExpectationsWereMet())

		})
	}
}

func TestRepository_CountIP(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	repo := NewCountryRepository(db)
	columns := []string{"country", "sum"}
	type TestCases struct {
		name           string
		countryCode    string
		expectedOutput IPCount
		expectedError  error
		scanError      bool
		closeError     bool
		rowError       bool
	}

	sampleIPCount := IPCount{
		Country:  "US",
		Quantity: 5,
	}

	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnIPCount",
			countryCode:    sampleIPCount.Country,
			expectedOutput: sampleIPCount,
			expectedError:  nil,
		},
		{
			name:           "DBError_ReturnError",
			countryCode:    sampleIPCount.Country,
			expectedOutput: IPCount{},
			expectedError:  errors.New("unexpected error"),
		},
		{
			name:           "ScanError_ReturnError",
			countryCode:    sampleIPCount.Country,
			expectedOutput: IPCount{},
			expectedError:  nil,
			scanError:      true,
		},
		{
			name:           "CloseError_ReturnError",
			countryCode:    sampleIPCount.Country,
			expectedOutput: IPCount{},
			expectedError:  nil,
			closeError:     true,
		},
	} {

		t.Run(tt.name, func(t *testing.T) {
			rows := sqlmock.NewRows(columns)
			if tt.expectedOutput.Country != "" {
				rows.AddRow(tt.expectedOutput.Country, tt.expectedOutput.Quantity)
			}
			exec := mock.ExpectQuery(regexp.QuoteMeta(countIPsQuery)).WithArgs(tt.countryCode)

			if tt.closeError {
				rows.CloseError(sql.ErrConnDone)
			}

			if tt.scanError {
				rows.AddRow("example", "bad value")
			}

			if tt.expectedError != nil {
				exec.WillReturnError(tt.expectedError)
			}

			exec.WillReturnRows(rows)

			res, err := repo.CountIPs(context.Background(), tt.countryCode)
			if tt.expectedError != nil || tt.rowError || tt.scanError || tt.closeError {
				assert.Empty(t, res)
				assert.Error(t, err)
				return
			}
			assert.NotNil(t, res)
			assert.Equal(t, tt.expectedOutput, res)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetIP(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	columns := []string{"ip_from", "ip_to", "proxy_type", "country_code", "country_name", "region_name", "city_name", "isp", "domain", "usage_type", "asn", "asn"}

	repo := NewCountryRepository(db)

	type TestCases struct {
		name           string
		ip             IPAddress
		expectedOutput IPModel
		expectedError  error
		scanError      bool
		closeError     bool
		rowError       bool
	}

	ipAddr, err := NewIPAddressFromString("111.111.1.11")
	assert.Nil(t, err)

	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnIPModel",
			ip:             ipAddr,
			expectedOutput: sampleIPModel(ipAddr, false),
			expectedError:  nil,
		},
		{
			name:           "InputHasDifferentIPs_ReturnIPModel",
			ip:             ipAddr,
			expectedOutput: sampleIPModel(ipAddr, true),
			expectedError:  nil,
		},
		{
			name:           "DBError_ReturnError",
			ip:             ipAddr,
			expectedOutput: IPModel{},
			expectedError:  errors.New("unexpected error"),
		},
		{
			name:           "ScanError_ReturnError",
			ip:             ipAddr,
			expectedOutput: IPModel{},
			expectedError:  nil,
			scanError:      true,
		},
		{
			name:           "CloseError_ReturnError",
			ip:             ipAddr,
			expectedOutput: IPModel{},
			expectedError:  nil,
			closeError:     true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			rows := sqlmock.NewRows(columns)

			if tt.expectedOutput.IPFrom > 0 {
				rows.AddRow(
					tt.expectedOutput.IPFrom,
					tt.expectedOutput.IPTo,
					tt.expectedOutput.ProxyType,
					tt.expectedOutput.CountryCode,
					tt.expectedOutput.CountryName,
					tt.expectedOutput.RegionName,
					tt.expectedOutput.CityName,
					tt.expectedOutput.ISP,
					tt.expectedOutput.Domain,
					tt.expectedOutput.UsageType,
					tt.expectedOutput.ASN,
					tt.expectedOutput.AS,
				)
			}

			exec := mock.ExpectQuery(regexp.QuoteMeta(getIPInfoQuery)).WithArgs(tt.ip.IntIP)

			if tt.closeError {
				rows.CloseError(sql.ErrConnDone)
			}

			if tt.scanError {
				rows.AddRow("bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad")
			}

			if tt.expectedError != nil {
				exec.WillReturnError(tt.expectedError)
			}

			exec.WillReturnRows(rows)

			res, err := repo.GetIP(context.Background(), tt.ip.IntIP)
			if tt.expectedError != nil || tt.rowError || tt.scanError || tt.closeError {
				assert.Empty(t, res)
				assert.Error(t, err)
				return
			}
			assert.NotNil(t, res)
			assert.Equal(t, tt.expectedOutput, res)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
	}
}
