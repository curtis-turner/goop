package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/curtis-turner/ics"
)

func main() {
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "confluent.customer.divvycloud.com",
	}
	client := ics.NewClient(baseURL, ics.WithCredentials("username", "password"))

	if err := client.NewSession(); err != nil {
		fmt.Printf("Error Occured: %v", err)
	}

	ctx := context.Background()

	res, _ := client.Resources().Get(ctx, "test")

	fmt.Printf("Resource Response: %v", res)

	resp, _ := client.Resources().List(ctx, ics.Limit(10))

	for _, resource := range resp.Resources {
		fmt.Printf("Resource: %v", resource)
	}
}
