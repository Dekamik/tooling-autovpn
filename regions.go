package main

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
)

type region struct {
    Id      string `json:"id"`
    Country string `json:"country"`
}

func getRegions() ([]region, error) {
    fmt.Print("Downloading regions...")
    defer fmt.Println("OK")
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

func showRegions() error {
    regions, regionsErr := getRegions()
    if regionsErr != nil {
        return regionsErr
    }

    var str = ""
    if options.PrintJson {
        jsonBytes, jsonErr := json.Marshal(regions)
        if jsonErr != nil {
            return jsonErr
        }
        str = string(jsonBytes)
        fmt.Println(str)
    } else {
        matrix := make([][]string, len(regions))
        for i, r := range regions {
            matrix[i] = []string { r.Id, r.Country }
        }
        if !options.NoHeaders {
            matrix = append([][]string{{"Region ID", "Country Code"}}, matrix...)
        }
        printTable(matrix)
    }

    return nil
}

func isRegionValid(regionName string) (bool, error) {
    regions, regionsErr := getRegions()
    if regionsErr != nil {
        return false, regionsErr
    }

    for _, region := range regions {
        if region.Id == regionName {
            return true, nil
        }
    }

    return false, nil
}
