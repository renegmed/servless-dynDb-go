package test

import (
	"fmt"
	"io/ioutil"
	util "serverless-movies-pjohnson/test_utils/utils"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemService_ListByYear(t *testing.T) {
	content, err := ioutil.ReadFile("../../data/post1.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Fatal("File1 is empty")
	}

	content2, err := ioutil.ReadFile("../../data/post2.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Fatal("File2 is empty")
	}

	service, err := util.NewItemService()
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.Post(string(content))
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.Post(string(content2))
	if err != nil {
		t.Fatal(err)
	}

	items, err := service.ListByYear("2013")
	if err != nil {
		fmt.Printf("Test error found:\n %v\n", err)
		assert.Error(t, err)
	} else {
		assert.Equal(t, 2, len(items))
		assert.Equal(t, 2013, items[0].YearReleased)
		assert.Equal(t, 2013, items[1].YearReleased)

		counter := 0
		for _, item := range items {
			if item.Title == "Turn It Down, Or Else!" {
				assert.Equal(t, "Turn It Down, Or Else!", items[counter].Title)
				goto valid
			}
			counter++
		}
		assert.Fail(t, "Failed to find title")

	valid:
	}

}

func TestItemService_GetByTitle(t *testing.T) {
	content, err := ioutil.ReadFile("../../data/post1.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Fatal("File1 is empty")
	}

	content2, err := ioutil.ReadFile("../../data/post2.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Fatal("File2 is empty")
	}

	service, err := util.NewItemService()
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.Post(string(content))
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.Post(string(content2))
	if err != nil {
		t.Fatal(err)
	}

	items, err := service.GetByTitle("Hunger Games: Catching Fire")
	if err != nil {
		fmt.Printf("Test error found:\n %v\n", err)
		assert.Error(t, err)
	} else {
		assert.Equal(t, 2013, items[0].YearReleased)
		assert.Equal(t, "Hunger Games: Catching Fire", items[0].Title)
	}

}
