package user

import "testing"

func TestBaseUserName(t *testing.T) {
	user := BaseUser{Username: "test"}
	if user.Name() != "test" {
		t.Errorf("user:BaseUser:Name returned the wrong name")
	}
}
