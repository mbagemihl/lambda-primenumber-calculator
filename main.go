package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"math"
	"strconv"
	"time"
)


type MyEvent struct {
	Maximum int `json:"max"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	start := time.Now()
	max, _ := strconv.Atoi(request.QueryStringParameters["max"])
	var result []int
	for i :=1; i < max; i++ {
		if prime := IsPrime(i); prime == true {
			result = append(result, i)
		}
	}
	fmt.Println(result)
	end := time.Now()
	fmt.Println(end.Sub(start))

	r, _ := json.Marshal(result)
	return events.APIGatewayProxyResponse{Body: string(r), StatusCode: 200}, nil
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}