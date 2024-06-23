package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Rate struct {
    TimePeriodStart string  `json:"time_period_start"`
    TimeClose       string  `json:"time_close"`
    RateOpen        float64 `json:"rate_open"`
    RateClose       float64 `json:"rate_close"`
}

func main() {
    // Set your API key here
    headers := map[string]string{
        "X-CoinAPI-Key": "YOUR_API_KEY",
    }
    //indicate the currencies you need
    lst := []string{"BTC", "ETH", "DOT", "SOL", "XRP", "LTC", "BNB", "TRX", "MATIC", "ADA"}
    params := map[string]string{
        "period_id":  "1DAY",
        "time_start": "2023-06-21T18:39:15.645Z",
        "time_end":   "2023-06-27T18:39:15.645Z",
    }

    for _, i := range lst {
        url := fmt.Sprintf("https://rest.coinapi.io/v1/exchangerate/%s/USD/history?period_id=%s&time_start=%s&time_end=%s",
            i, params["period_id"], params["time_start"], params["time_end"])

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            fmt.Println("Error creating request:", err)
            continue
        }

        for key, value := range headers {
            req.Header.Set(key, value)
        }

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            fmt.Println("Error making request:", err)
            continue
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response body:", err)
            continue
        }

        var rates []Rate
        err = json.Unmarshal(body, &rates)
        if err != nil {
            fmt.Println("Error unmarshaling response:", err)
            continue
        }

        fmt.Printf("Name: %s\n", i)
        fmt.Println(" Day   Opening and closing times     Opening and closing rate")

        for index, rate := range rates {
            fmt.Printf("   %d      %s  %.2f$\n", index+1, rate.TimePeriodStart, rate.RateOpen)
            fmt.Printf("            %s  %.2f$\n", rate.TimeClose, rate.RateClose)
        }

        fmt.Println()
    }
}
