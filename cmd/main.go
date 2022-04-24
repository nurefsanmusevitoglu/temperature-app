package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	customerror "github.com/nurefsanmusevitoglu/temperature-app/pkg/error"
	"github.com/nurefsanmusevitoglu/temperature-app/pkg/weather"
)

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		switch req.Resource {
		case "/temperature":
			return getTemperature(req)
		default:
			log.Println("path does not exist...")
			return customerror.New(http.StatusNotFound), nil
		}
	default:
		return customerror.New(http.StatusMethodNotAllowed), nil
	}
}

func getTemperature(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var temp weather.Temperature

	city := req.QueryStringParameters["city"]
	if city == "" {
		log.Println("city is missing...")
		return customerror.New(http.StatusBadRequest), nil
	}

	if os.Getenv("apikey") == "" {
		log.Println("api.openweathermap.org apikey is required...")
		return customerror.New(http.StatusBadRequest), nil
	}

	temp.GetTemperature(city)
	if temp.IsEmpty() {
		return customerror.New(http.StatusBadRequest), nil
	}

	response, err := json.Marshal(temp)
	if err != nil {
		log.Println("temp info cannot be marshalled...")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(response),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
