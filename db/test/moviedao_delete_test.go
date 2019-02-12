package test

import (
	"fmt"
	"io/ioutil"
	util "serverless-movies-pjohnson/test_utils/utils"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemService_DeleteById(t *testing.T) {
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

	var idToDelete string

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
				idToDelete = items[counter].ID
				goto valid
			}
			counter++
		}
		assert.Fail(t, "Failed to find title")

	valid:
	}

	oldItem, err := service.Delete(idToDelete)
	if err != nil {
		fmt.Printf("Test error found:\n %v\n", err)
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		fmt.Printf("Deleted old item ID:\n %s\n", oldItem.ID)
		fmt.Printf("Deleted old item title:\n %s\n", oldItem.Title)
		fmt.Printf("Deleted old item year: %d\n", oldItem.YearReleased)

		assert.Equal(t, 2013, oldItem.YearReleased)
		assert.Equal(t, "Turn It Down, Or Else!", oldItem.Title)
	}
}

func TestItemService_DeleteByYearTitle(t *testing.T) {
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

	//fmt.Printf("Content: \n %s\n", string(content))

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

	oldItem, err := service.DeleteByYearTitle("2013", "Turn It Down, Or Else!")
	if err != nil {
		fmt.Printf("Test error found:\n %v\n", err)
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, 2013, oldItem.YearReleased)
		assert.Equal(t, "Turn It Down, Or Else!", oldItem.Title)
	}

	items, err := service.ListByYear("2013")
	if err != nil {
		fmt.Printf("Test error found:\n %v\n", err)
		assert.Error(t, err)
	} else {
		assert.Equal(t, 1, len(items))
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
