package main

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Username string  `json:"username,omitempty"`
	Id       float64 `json:"id,omitempty"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if request.HTTPMethod == "POST" {

		var jsonResponse Request
		json.Unmarshal([]byte(request.Body), &jsonResponse)

		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("{message: Hello, %s! Your Id is %f}", jsonResponse.Username, jsonResponse.Id),
			StatusCode: 200,
		}, nil

	} else if request.HTTPMethod == "GET" {

		return events.APIGatewayProxyResponse{
			Body:       "{message: Hello, World!}",
			StatusCode: 200,
		}, nil

	} else {

		return events.APIGatewayProxyResponse{
			Body:       "Ooops",
			StatusCode: 502,
		}, errors.New("HTTP Method not allowed!")

	}
}

func main() {
	lambda.Start(Handler)
}
