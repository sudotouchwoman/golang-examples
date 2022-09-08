package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type webAPI struct {
	host    string
	version string
}

func (api *webAPI) fetchQuotes(num uint) Quotes {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/%s/quotes/%d", api.host, api.version, num),
		nil,
	)

	// abort the request if it takes too long (more than 5s)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	// try to parse obtained json into a struct
	var qs Quotes
	if err := json.NewDecoder(resp.Body).Decode(&qs); err != nil {
		panic(err)
	}
	return qs
}

func (api *webAPI) fetchQuote() Quote {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/%s/quotes", api.host, api.version),
		nil,
	)

	// abort the request if it takes too long (more than 5s)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	// try to parse obtained json into a struct
	var q Quotes
	if err := json.NewDecoder(resp.Body).Decode(&q); err != nil || len(q) != 1 {
		panic(err)
	}
	return q[0]
}

var quotesAPI webAPI

func init() {
	if err := godotenv.Load(); err != nil {
		log.Default().Fatal("Error loading .env file")
	}

	quotesAPI.version = os.Getenv("QUOTES_API_VERSION")
	quotesAPI.host = os.Getenv("QUOTES_API_HOST")

	log.Default().
		Printf("Loaded API info: %s/%s", quotesAPI.host, quotesAPI.version)
}
