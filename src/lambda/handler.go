package main

import (
	"github.com/asragi/spajam2023lambda/greeting"
	"github.com/aws/aws-lambda-go/lambda"
)

func greet() {
	greeting.Hello()
}

func main() {
	lambda.Start(greet)
}
