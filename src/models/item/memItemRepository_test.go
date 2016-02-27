package item

import "testing"

func TestNewItem(t *testing.T) {
	var memItemRepo memItemRepository
	newItem, err := memItemRepo.NewItem("test", "item")
	if err != nil {
		t.Errorf(err.Error())
	}
	if newItem.Name() != "test" {
		t.Errorf("Created Item did not set the name properly")
	}
}

func TestGetItemById(t *testing.T) {
	var memItemRepo memItemRepository
	newItem, err := memItemRepo.NewItem("test", "item")
	if err != nil {
		t.Errorf(err.Error())
	}
	gotItem, err := memItemRepo.GetItemById(newItem.Id())
	if gotItem.Name() != "test" {
		t.Errorf("Item retrieved by Id had the incorrect name")
	}
}
