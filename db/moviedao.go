package db

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/satori/go.uuid"
)

// ItemInfo has more data for our movie item
type ItemInfo struct {
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}

// Item has fields for the DynamoDB keys (Year and Title) and an ItemInfo for more data
type Item struct {
	ID           string   `json:"id" dynamo:"ID,hash"`
	Title        string   `json:"title" dynamo:"Title"`
	YearReleased int      `json:"yearReleased" dynamo:"YearReleased"`
	Info         ItemInfo `json:"info" dynamo:"Info"`
}

// MoviesService holds out dynamo client
type ItemService struct {
	Table dynamo.Table
}

// NewItemService creates a new item service with a dynamo client setup to talk to the provided table name
func NewItemService() (*ItemService, error) {
	dynamoTable, err := NewDynamoTable(os.Getenv("TABLE_NAME"), "")
	if err != nil {
		return nil, err
	}
	return &ItemService{
		Table: dynamoTable,
	}, nil
}

func NewDynamoTable(tableName, endpoint string) (dynamo.Table, error) {
	if tableName == "" {
		return dynamo.Table{}, fmt.Errorf("you must supply a table name")
	}
	cfg := aws.Config{}
	if endpoint != "" {
		cfg.Region = aws.String("us-east-1")
		cfg.Endpoint = aws.String(endpoint)
	}

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &cfg)
	table := db.Table(tableName)
	return table, nil
}

func (i *ItemService) GetById(id string) (Item, error) {
	var result Item

	err := i.Table.Get("ID", id).Consistent(true).One(&result)

	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}

	return result, nil
}

func (i *ItemService) GetByTitle(title string) ([]Item, error) {
	var items []Item

	//err := i.Table.Get("YearReleased", yearAsInt).Consistent(true).Filter("Title = ?", title).One(&result)
	//err := i.Table.Get("Title", title).Consistent(true).One(&result)
	err := i.Table.Scan().Filter("Title = ?", title).All(&items)
	if err != nil {
		fmt.Println(err.Error())
		return items, err
	}

	return items, nil
}

func (i *ItemService) ListByYear(year string) ([]Item, error) {

	yearAsInt, _ := strconv.Atoi(year)

	var items []Item
	err := i.Table.Scan().Filter("'YearReleased' = ?", yearAsInt).All(&items)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return items, nil
}

func (i *ItemService) Post(body string) (Item, error) {

	var thisItem Item
	json.Unmarshal([]byte(body), &thisItem)

	id, err := uuid.NewV4()
	if err != nil {
		return thisItem, err
	}
	thisItem.ID = id.String()
	i.Table.Put(thisItem).Run()
	return thisItem, nil
}

func (i *ItemService) Delete(id string) (Item, error) {

	var oldItem Item

	err := i.Table.Delete("ID", id).OldValue(&oldItem)
	if err != nil {
		fmt.Println(err.Error())
		return oldItem, err
	}
	return oldItem, nil
}

// This doesn't work
//    ConditionalCheckFailedException: The conditional request failed
func (i *ItemService) DeleteByYearTitle(year, title string) (Item, error) {

	yearAsInt, _ := strconv.Atoi(year)
	var oldItem Item

	err := i.Table.Delete("ID", "*").
		If("YearReleased = ? AND Title = ?", yearAsInt, title).
		OldValue(&oldItem)

	if err != nil {
		fmt.Println(err.Error())
		return oldItem, err
	}

	fmt.Printf("++++++Old Item deleted: \n   %v \n", oldItem)

	return oldItem, nil
}

func (i *ItemService) Put(newItem Item) (Item, error) {

	var oldItem Item
	var cc dynamo.ConsumedCapacity
	err := i.Table.Put(newItem).ConsumedCapacity(&cc).OldValue(&oldItem)

	// check for putting the same item: this should fail
	// err := i.Table.Put(newItem).If("attribute_not_exists(Title)").Run()
	if err != nil {
		return oldItem, err
	}

	if cc.Total != 1 || cc.Table != 1 { //|| cc.TableName != testTable {
		return oldItem, fmt.Errorf("bad consumed capacity: %#v", cc)
	}

	return oldItem, nil
}
