# Cryptocurrency Exchange Rate History

This Go program fetches historical exchange rate data for selected cryptocurrencies against USD using the CoinAPI service. The program retrieves the data for a specified time period and prints the opening and closing rates for each day within that period.

## Features

- Fetches historical exchange rate data for a list of cryptocurrencies.
- Prints the opening and closing times along with their respective rates.
- Easy to modify to include other cryptocurrencies or time periods.

## Prerequisites

- Go 1.16 or later

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/crypto-exchange-rate-history.git
    cd crypto-exchange-rate-history
    ```

2. Get dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Replace the `X-CoinAPI-Key` in the code with your own CoinAPI key. You can get one by signing up at [CoinAPI](https://www.coinapi.io/).

2. Run the program:

    ```sh
    go run main.go
    ```

## Customization

- To fetch data for different cryptocurrencies, modify the `currencies` slice in the `main` function:

    ```go
    currencies := []string{"BTC", "ETH", "DOT", "SOL", "XRP", "LTC", "BNB", "TRX", "MATIC", "ADA"}
    ```

- To change the time period, update the `params` map:

    ```go
    params := map[string]string{
        "period_id":  "1DAY",
        "time_start": "2023-06-21T18:39:15.645Z",
        "time_end":   "2023-06-27T18:39:15.645Z",
    }
    ```

## Example Output

