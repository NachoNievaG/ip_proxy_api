package country

type ISPCount struct {
	ISP      string `json:"isp"`
	Quantity int64  `json:"quantity"`
}

type IPCount struct {
	Country  string `json:"country"`
	Quantity int64  `json:"quantity"`
}

type IPModel struct {
	IPFrom      int64  `json:"-"`
	IPTo        int64  `json:"-"`
	ProxyType   string `json:"proxy_type"`
	CountryCode string `json:"countr_code"`
	CountryName string `json:"countr_name"`
	RegionName  string `json:"region_name"`
	CityName    string `json:"city_name"`
	ISP         string `json:"isp"`
	Domain      string `json:"domain"`
	UsageType   string `json:"usage_type"`
	ASN         string `json:"asn"`
	AS          string `json:"as"`
}

type IPResponse struct {
	IPFrom string `json:"ip_from"`
	IPTo   string `json:"ip_to"`
	IPModel
}

func (m IPModel) ToResponse() IPResponse {
	ipFrom := NewIPAddressFromBigEndian(uint32(m.IPFrom))
	ipTo := NewIPAddressFromBigEndian(uint32(m.IPTo))
	return IPResponse{
		IPFrom:  ipFrom.IP.String(),
		IPTo:    ipTo.IP.String(),
		IPModel: m,
	}
}
