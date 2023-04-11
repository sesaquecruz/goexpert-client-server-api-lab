package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const RequestTimeout = 300 * time.Millisecond

type Quote struct {
	Bid float64 `json:"bid,string"`
}

func main() {
	url := "http://localhost:8080/cotacao"

	// Get quote from server
	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Fail - status code: %d\n", res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Panic(err)
	}

	// Show quote
	log.Printf("Quote - %+v\n", quote)

	// Save quote
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Dólar: {%f}\n", quote.Bid)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	log.Println("The quote was saved")
}
