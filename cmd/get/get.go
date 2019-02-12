package main

import (
	"encoding/json"
	"fmt"
	"serverless-movies-pjohnson/db"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

// Parse slug into a space separated string
func parseSlug(orig string) (retval string) {
	retval = strings.Replace(orig, "-", " ", -1)
	retval = strings.Replace(retval, "+", " ", -1)
	return retval
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Make the call to the DAO with params found in the path
	fmt.Println("Path vars: ", request.PathParameters["title"])

	db, err := db.NewItemService()
	if err != nil {
		panic(fmt.Sprintf("Get: Failed to connect to table:\n %v", err))
	}

	items, err := db.GetByTitle(request.PathParameters["title"])
	if err != nil {
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}

	// Make sure the Item isn't empty
	if len(items) > 0 {
		fmt.Println("Could not find movie")
		return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 500}, nil
	}

	// Log and return result
	jsonItem, _ := json.Marshal(items)
	stringItem := string(jsonItem) + "\n"
	fmt.Println("Found item: ", stringItem)
	return events.APIGatewayProxyResponse{Body: stringItem, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
