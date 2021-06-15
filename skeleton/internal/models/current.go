package models

import "github.com/ychengcloud/api/internal/generated/models"

type CurrentUser struct {
	Name        string
	Permissions []Permission
}

type Permission struct {
	Code string
	Name string
}
