package main

import (
	"context"
	"net/http"
	"time"

	"github.com/brittonhayes/aos/client"
)

func main() {
	//	Setup a context with a timeout so we're covered in case of a slow response
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new http client
	c := client.NewClient(&http.Client{}, "https://aos-api.com/query", nil)

	// Get all allegiances
	resp, err := c.GetAllegiances(ctx, client.AllegianceFilters{})
	if err != nil {
		panic(err)
	}

	// List the allegiances
	for _, a := range resp.Allegiances {
		println(a.Name)
	}
}
