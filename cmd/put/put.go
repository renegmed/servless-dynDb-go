package main

import (
	"encoding/json"
	"fmt"
	"serverless-movies-pjohnson/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Log body and pass to the DAO
	fmt.Println("Received body: ", request.Body)

	var item db.Item
	json.Unmarshal([]byte(request.Body), &item)

	db, err := db.NewItemService()
	if err != nil {
		panic(fmt.Sprintf("Put: Failed to connect to table:\n %v", err))
	}

	updatedItem, err := db.Put(item)
	if err != nil {
		fmt.Println("Got error calling put")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 500}, nil
	}

	// Log and return result
	fmt.Printf("Updated item: %v \n ", updatedItem)
	return events.APIGatewayProxyResponse{Body: "Success\n", StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
