package tooling

import (
	"fmt"
	"net/http"
)

type RespLocation struct {
	Status     string `json:"status"`
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
	City       string `json:"city"`
	Timezone   string `json:"timezone"`
}

func GetLocation(ip string) (string, string, string, error) {
	var respLocation RespLocation

	Response, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,country,regionName,city,timezone", ip))
	if err != nil {
		return "", "", "", err
	}

	err = UnmarshallAll(Response.Body, respLocation)
	if err != nil {
		return "", "", "", err
	}

	return respLocation.Country, respLocation.RegionName, respLocation.City, nil
}
