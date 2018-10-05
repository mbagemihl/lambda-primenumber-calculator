package main

import (
	"awesomeFunction/calc"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	max, _ := strconv.Atoi(request.QueryStringParameters["max"])
	result := calc.GetPrimeNumbersBelowMax(max)
	r, _ := json.Marshal(result)
	return events.APIGatewayProxyResponse{Body: string(r), StatusCode: 200}, nil
}