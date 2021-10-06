package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type region struct {
	Id string `json:"id"`
	Country string `json:"country"`
}

func getRegions() ([]region, error) {
	type regionRes struct {
		Data []region `json:"data"`
	}

	url := "https://api.linode.com/v4/regions"
	client := http.Client{}
	req, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		return nil, requestErr
	}
	req.Header.Set("User-Agent", "Dekamik/autovpn")

	res, getErr := client.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(res.Body)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	regions := regionRes{}
	jsonErr := json.Unmarshal(body, &regions)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return regions.Data, nil
}

func isRegion(str string, regions []region) bool {
	for _, region := range regions {
		if str == region.Id {
			return true
		}
	}
	return false
}

func showRegions() error {
	verbose("Fetching regions...")
	regions, regionsErr := getRegions()
	if regionsErr != nil {
		return regionsErr
	}
	verboseln("OK")

	var str = ""
	if options.PrintJson {
		jsonBytes, jsonErr := json.Marshal(regions)
		if jsonErr != nil {
			return jsonErr
		}
		str = string(jsonBytes)
		fmt.Println(str)
	} else {
		matrixExample := [][]interface{} {
			{"A", 1},
			{"B", 2},
		}
		fmt.Println(matrixExample)
		matrix := Map(regions, func(item interface{}) interface{} { return []interface{} { item.(region).Id, item.(region).Country } })
		fmt.Println(matrix)
		printMatrix(matrix)
		if options.PrintHeaders {
			regions = append([]region{{Id: "Region ID", Country: "Country Code"}}, regions...)
		}
		format := fmt.Sprintf("%%-%ds %%-%ds\n",
			maxlen(Map(regions, func(item interface{}) interface{} { return item.(region).Id })),
			maxlen(Map(regions, func(item interface{}) interface{} { return item.(region).Country })))
		for _, region := range regions {
			fmt.Printf(format, region.Id, region.Country)
		}
	}

	return nil
}

func validateRegions(regionArgs []string) error {
	if len(regionArgs) == 0 {
		return errors.New("No region specified ")
	}

	verbose("Fetching regions...")
	regions, regionsErr := getRegions()
	if regionsErr != nil {
		return regionsErr
	}
	verboseln("OK")

	for _, region := range regionArgs {
		if !isRegion(region, regions) {
			return errors.New(fmt.Sprintf("Illegal region \"%s\"", region))
		}
	}

	return nil
}
