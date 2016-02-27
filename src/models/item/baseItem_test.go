package item

import "testing"

func TestBaseItemName(t *testing.T) {
	item := BaseItem{name: "test", id: "zero"}
	if item.Name() != "test" {
		t.Errorf("item:BaseItem:Name returned the wrong name")
	}
}

func TestBaseItemId(t *testing.T) {
	item := BaseItem{name: "test", id: "12345"}
	if item.Id() != "12345" {
		t.Errorf("item:BaseItem:Id returned the wrong Id")
	}
}

func TestBaseItemSetName(t *testing.T) {
	item := BaseItem{name: "test", id: "12345"}
	item.SetName("newName")
	if item.Name() != "newName" {
		t.Errorf("item:BaseItem:SetName returned the wrong new name")
	}
}
