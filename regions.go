package main

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
)

type Region struct {
    Id      string `json:"id"`
    Country string `json:"country"`
}

func getRegions() ([]Region, error) {
    fmt.Print("Downloading regions...")
    defer fmt.Println("OK")
    type regionRes struct {
        Data []Region `json:"data"`
    }

    url := "https://api.linode.com/v4/regions"
    client := http.Client{}
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil { return nil, err }
    req.Header.Set("User-Agent", "Dekamik/autovpn")

    res, err := client.Do(req)
    if err != nil { return nil, err }
    if res.Body != nil {
        defer func(Body io.ReadCloser) {
            err = Body.Close()
            if err != nil {
                panic(err)
            }
        }(res.Body)
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil { return nil, err }

    regions := regionRes{}
    err = json.Unmarshal(body, &regions)
    if err != nil { return nil, err }

    return regions.Data, nil
}

func showRegions() error {
    regions, err := getRegions()
    if err != nil { return err }

    var str = ""
    if options.PrintJson {
        jsonBytes, err := json.Marshal(regions)
        if err != nil { return err }
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
    regions, err := getRegions()
    if err != nil { return false, err }

    for _, region := range regions {
        if region.Id == regionName {
            return true, nil
        }
    }

    return false, nil
}
