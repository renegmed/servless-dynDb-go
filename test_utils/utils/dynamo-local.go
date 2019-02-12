package util

import (
	"fmt"
	"serverless-movies-pjohnson/db"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/rs/xid"
)

func NewItemService() (*db.ItemService, error) {

	tableName, err := CreateItemTable(db.Item{})
	if err != nil {
		return nil, fmt.Errorf("failed to set up table. %s", err)
	}

	itemdb, err := db.NewDynamoTable(tableName, "http://localhost:9000") // docker endpoint
	if err != nil {
		return nil, err
	}

	service := db.ItemService{
		Table: itemdb,
	}

	return &service, nil
}

func CreateItemTable(table interface{}) (string, error) {
	cfg := aws.Config{
		Endpoint: aws.String("http://localhost:9000"),
		Region:   aws.String("us-east-1"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	}

	sess := session.Must(session.NewSession())

	db := dynamo.New(sess, &cfg)
	tableName := xid.New().String()

	err := db.CreateTable(tableName, table).Run()
	if err != nil {
		return "", err
	}

	return tableName, nil
}
