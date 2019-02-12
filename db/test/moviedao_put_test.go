package test

import (
	"fmt"
	"io/ioutil"

	"testing"

	util "serverless-movies-pjohnson/test_utils/utils"

	"github.com/stretchr/testify/assert"
)

func TestItemService_UpdateItem(t *testing.T) {

	content, err := ioutil.ReadFile("../../data/post1.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Fatal("File is empty")
	}

	// content2, err := ioutil.ReadFile("../../data/post2.json")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if len(content) == 0 {
	// 	t.Fatal("File2 is empty")
	// }

	//fmt.Printf("Content: \n %s\n", string(content))

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

	// item2, err := service.Post(string(content2))
	// if err != nil {
	// 	assert.Error(t, err)
	// } else {
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, 2013, item2.YearReleased)
	// 	assert.Equal(t, "Hunger Games: Catching Fire", item2.Title)
	// }

	// +++++++ change title here  here for content 1 ++++++++

	// content, err = ioutil.ReadFile("../../data/post1_put.json")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if len(content) == 0 {
	// 	t.Fatal("File post1_put.json is empty")
	// }

	// make a change here
	item.YearReleased = 2014
	oldItem, err := service.Put(item)
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, 2013, oldItem.YearReleased)
		assert.Equal(t, "Turn It Down, Or Else!", oldItem.Title)
	}

	items, err := service.GetByTitle("Turn It Down, Or Else!")
	if err != nil {
		fmt.Printf("Test error found:\n %v\n", err)
		assert.Error(t, err)
	} else {
		assert.Equal(t, 1, len(items))
		assert.Equal(t, 2014, items[0].YearReleased)
		assert.Equal(t, "Turn It Down, Or Else!", items[0].Title)
	}

}
