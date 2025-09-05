package model

import (
	"testing"
)

func TestRoleConstants(t *testing.T) {
	if RoleAdmin != "admin" {
		t.Errorf("Expected RoleAdmin to be \"admin\", got %q", RoleAdmin)
	}
	if RoleUser != "user" {
		t.Errorf("Expected RoleUser to be \"user\", got %q", RoleUser)
	}
}
