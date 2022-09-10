# ip_proxy_api
DreamLab Challenge Implementation
## Summary
IP proxy api is an API requirement impulsed by DreamLab Technologies AG.
The definition of this API consts of 3 main endpoints:
```
	Get /countries/CH/top_ten_isp
	Get /countries/{countryCode}/ip/count
	Get /ip/{ip}
```

### - Get /countries/CH/top_ten_isp

This service will provide the requester the top ten internet service providers from Switzerland, in descending order.
The response in case of success should be as this example:
```
[
  {
    "isp": "ISP 1",
    "quantity": 10
  },
  {
    "isp": "ISP 2",
    "quantity": 9
  },
  {
    "isp": "ISP 3",
    "quantity": 8
  },
  {
    "isp": "ISP 4",
    "quantity": 7
  },
  {
    "isp": "ISP 5",
    "quantity": 6
  },
  {
    "isp": "ISP 6",
    "quantity": 5
  },
  {
    "isp": "ISP 7",
    "quantity": 4
  },
  {
    "isp": "ISP 8",
    "quantity": 3
  },
  {
    "isp": "ISP 9",
    "quantity": 2
  },
  {
    "isp": "ISP 10",
    "quantity": 1
  }
]
```
_Note: this endpoint was developed having in mind it could be refactored to support other country codes_


### -	Get /countries/{countryCode}/ip/count
Using the country code, count every ip given in the database. In case the count is 0 (or unexistent country), given that is not a straight get to a specific property of a database entity, `0` is provided as count result.

The response in case of success should be as this example:
```
{
  "country": "AR",
  "quantity": 10
}
```
In case the count turns to be 0:
```
{
  "country": "AR",
  "quantity": 0
}
```

### - Get /ip/{ip}
Provided by an IP Address, return every field related to it.

The response in case of success should be as this example:
```
{
  "ip_from": "111.111.1.11",
  "ip_to": "111.111.1.11",
  "proxy_type": "PUB",
  "countr_code": "GB",
  "countr_name": "Great Britain",
  "region_name": "England",
  "city_name": "London",
  "isp": "sample ISP",
  "domain": "sample domain",
  "usage_type": "sample usage_type",
  "asn": "sample asn",
  "as": "sample as"
}
```

In case the provided IP is invalid:
```
{
  "status": 400,
  "result": "NewIPAddress: invalid ip"
}
```
In case the ip is not found:
```
{
  "status": 404,
  "result": "gtw.repository.GetIP: required information was not found"
}

```


