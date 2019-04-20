package main

import (
	"encoding/json"
	"fmt"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in go
// To trigger Number resource methods invoke corresponding helper method in main.go

const number = "1314315XXXX"

//Example to get details of number
//Pass number to fetch the details of a number

func TestNumberGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Numbers.Get(
		number,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}


//Example to update details of number
//Passappid of application to be assigned to the phone number, and NumberUpdateParams to update

func TestNumberUpdate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Numbers.Update(number,
		plivo.NumberUpdateParams{
			"app_id",
			"sub_account_id",
			"Updated Alias",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}


//Example to list all numbers
func TestNumberList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Numbers.List(
		plivo.NumberListParams{
			Limit: 5,
			Offset: 0,
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to delete a number
//Pass the number to delete
func TestNumberDelete(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Numbers.Delete(number)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted number successfully")
}

// Example to add number from your own carrier,
// Comma sperated numbers, carrier and region are mandatory params
func TestNumberCreate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Numbers.Create(
		plivo.NumberCreateParams{
			number,
			"carrier_id",
			"India",
			"number_type",
			"app_id",
			"sub_account_id",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

// Example to list all phone-number are available for purchase
// Country iso is mandatory param
func TestNumbersList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.PhoneNumbers.List(
		plivo.PhoneNumberListParams{
			CountryISO: "GB",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}


// Example to buy a phoneNumber
// Pass the numnber and app_id of application to be assigned to the phone number
func TestPhloneNumberCreate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.PhoneNumbers.Create(
		number,
		plivo.PhoneNumberCreateParams{
			"app_id",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

