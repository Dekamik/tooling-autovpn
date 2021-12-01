package main

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
)

type LinodeType struct {
    Id           string  `json:"id"`
    HourlyPrice  float32 `json:"price.hourly"`
    MonthlyPrice float32 `json:"price.monthly"`
}

func getTypes() ([]LinodeType, error) {
    fmt.Print("Downloading instance types...")
    defer fmt.Println("OK")
    type typeRes struct {
        Data []LinodeType `json:"data"`
    }

    url := "https://api.linode.com/v4/linode/types"
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

    types := typeRes{}
    err = json.Unmarshal(body, &types)
    if err != nil { return nil, err }

    return types.Data, nil
}

func showTypes() error {
    types, err := getTypes()
    if err != nil { return err }

    var str = ""
    if options.PrintJson {
        jsonBytes, err := json.Marshal(types)
        if err != nil { return err }
        str = string(jsonBytes)
        fmt.Println(str)
    } else {
        matrix := make([][]string, len(types))
        for i, r := range types {
            matrix[i] = []string { r.Id, fmt.Sprintf("%f", r.HourlyPrice), fmt.Sprintf("%f", r.MonthlyPrice) }
        }
        if !options.NoHeaders {
            matrix = append([][]string{{"ID", "Hourly price (USD)", "Monthly price (USD)"}}, matrix...)
        }
        printTable(matrix)
    }

    return nil
}

func isTypeValid(typeId string) (bool, error) {
    types, err := getTypes()
    if err != nil {
        return false, err
    }

    for _, linType := range types {
        if linType.Id == typeId {
            return true, nil
        }
    }

    return false, nil
}
