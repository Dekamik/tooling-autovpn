package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type region struct {
	Id string `json:"id"`
	Country string `json:"country"`
}

func getRegions() []region {
	type regionRes struct {
		Data []region `json:"data"`
	}

	url := "https://api.linode.com/v4/regions"
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	check(err)
	req.Header.Set("User-Agent", "Dekamik/autovpn")

	res, getErr := client.Do(req)
	check(getErr)
	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			check(err)
		}(res.Body)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	check(readErr)

	regions := regionRes{}
	jsonErr := json.Unmarshal(body, &regions)
	check(jsonErr)

	return regions.Data
}

func isRegion(str string, regions []region) bool {
	for _, region := range regions {
		if str == region.Id {
			return true
		}
	}
	return false
}
