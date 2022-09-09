package country

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGateway_GetTopTenISP(t *testing.T) {
	// Set controller mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := NewMockCountryRepository(mockCtrl)
	countryGateway := NewCountryGateway(m)

	type TestCases struct {
		name           string
		countryCode    string
		expectedOutput []ISPCount
		expectedError  error
	}

	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnTopTenISP",
			countryCode:    "CH",
			expectedOutput: sampleISP(),
			expectedError:  nil,
		},
		{
			name:           "RepoFails_ReturnError",
			countryCode:    "CH",
			expectedOutput: nil,
			expectedError:  errors.New("unexpected error"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			m.EXPECT().TopISPByCountryCode(ctx, "CH").Return(tt.expectedOutput, tt.expectedError).Times(1)

			res, err := countryGateway.TopTenISPByCountryCode(ctx, "CH")
			if tt.expectedError != nil {
				assert.Nil(t, res)
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, len(tt.expectedOutput), len(res))
			for i := range res {
				assert.Equal(t, tt.expectedOutput[i], res[i])
			}

		})
	}
}

func sampleISP() []ISPCount {
	var sampleISP []ISPCount
	for i := 1; i <= 10; i++ {
		sampleISP = append(sampleISP, ISPCount{"mocked internet provider", int64(i) * 10})
	}
	return sampleISP
}

func TestGateway_CountIP(t *testing.T) {
	// Set controller mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := NewMockCountryRepository(mockCtrl)
	countryGateway := NewCountryGateway(m)

	type TestCases struct {
		name           string
		countryCode    string
		expectedOutput IPCount
		expectedError  error
	}

	sampleIPCount := IPCount{
		Country:  "CH",
		Quantity: 5,
	}

	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnIPCount",
			countryCode:    "CH",
			expectedOutput: sampleIPCount,
			expectedError:  nil,
		},
		{
			name:           "RepoFails_ReturnError",
			countryCode:    "CH",
			expectedOutput: IPCount{},
			expectedError:  errors.New("unexpected error"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			m.EXPECT().CountIPs(ctx, "CH").Return(tt.expectedOutput, tt.expectedError).Times(1)

			res, err := countryGateway.CountIPsByCountryCode(ctx, "CH")
			if tt.expectedError != nil {
				assert.Empty(t, res)
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}

func TestGateway_GetIP(t *testing.T) {
	// Set controller mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := NewMockCountryRepository(mockCtrl)
	countryGateway := NewCountryGateway(m)

	stringIP := "192.111.1.11"
	testIP, err := NewIPAddressFromString(stringIP)
	assert.Nil(t, err)

	type TestCases struct {
		name           string
		ip             string
		expectedOutput IPModel
		expectedError  error
	}

	for _, tt := range []TestCases{
		{
			name:           "InputIsOK_ReturnIPInformation",
			ip:             stringIP,
			expectedOutput: sampleIPModel(testIP, false),
			expectedError:  nil,
		},
		{
			name:           "InputIsOK_ReturnIPInformation_DifferentIP",
			ip:             stringIP,
			expectedOutput: sampleIPModel(testIP, true),
			expectedError:  nil,
		},
		{
			name:           "RepoFails_ReturnError",
			ip:             stringIP,
			expectedOutput: IPModel{},
			expectedError:  errors.New("unexpected error"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			m.EXPECT().GetIP(ctx, testIP.IntIP).Return(tt.expectedOutput, tt.expectedError).Times(1)

			res, err := countryGateway.GetIPInfo(ctx, tt.ip)
			if tt.expectedError != nil {
				assert.Empty(t, res)
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.expectedOutput.ToResponse(), res)
		})
	}
}

func TestGetIP_IPIsInvalid_ReturnInvalidIPErr(t *testing.T) {
	countryGateway := NewCountryGateway(nil)
	badIP := "BADIPADDRESS"

	res, err := countryGateway.GetIPInfo(context.Background(), badIP)

	assert.Empty(t, res)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrInvalidIP))

}

func sampleIPModel(ip IPAddress, useDifferentIPs bool) IPModel {
	ipTo := ip.IntIP
	if useDifferentIPs {
		ipTo = ip.IntIP + 10
	}
	return IPModel{
		IPFrom:      int64(ip.IntIP),
		IPTo:        int64(ipTo),
		ProxyType:   "Test Proxy",
		CountryCode: "TC",
		CountryName: "Test Country",
		RegionName:  "Test Region",
		CityName:    "Test City",
		ISP:         "Test ISP",
		Domain:      "Test Domain",
		UsageType:   "Test Usage",
		ASN:         "Test ASN",
		AS:          "Test AS",
	}
}
