package main

import (
	"encoding/json"
	"fmt"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in credentials.go
// To trigger Pricing resource methods invoke corresponding helper method in main.go


// Example to get pricing
// Country ISO is mandatory param

func TestPricingGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Pricing.Get(
		"IN",
	)
	if err != nil {
		panic(err)
	}

	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}