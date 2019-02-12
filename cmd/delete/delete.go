package main

import (
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
	fmt.Println("Path vars: ", request.PathParameters["id"])

	db, err := db.NewItemService()
	if err != nil {
		panic(fmt.Sprintf("Delete: Failed to connect to table:\n %v", err))
	}

	//err = db.Delete(request.PathParameters["year"], parseSlug(request.PathParameters["title"]))
	_, err = db.Delete(request.PathParameters["id"])
	if err != nil {
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}
	return events.APIGatewayProxyResponse{Body: "Success\n", StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
