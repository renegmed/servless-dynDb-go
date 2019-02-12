package test

import (
	"io/ioutil"

	"testing"

	util "serverless-movies-pjohnson/test_utils/utils"

	"github.com/stretchr/testify/assert"
)

func TestItemService_CreateItem(t *testing.T) {

	content, err := ioutil.ReadFile("../../data/post1.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Fatal("File is empty")
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

	item, err := service.Post(string(content))
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, 2013, item.YearReleased)
		assert.Equal(t, "Turn It Down, Or Else!", item.Title)
	}

	item2, err := service.Post(string(content2))
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, 2013, item2.YearReleased)
		assert.Equal(t, "Hunger Games: Catching Fire", item2.Title)
	}
}
