package model

import (
	"github.com/MisakaTAT/GTerm/backend/enums"
)

type Credential struct {
	Common
	Name               string         `json:"name" gorm:"uniqueIndex;not null"`
	Username           string         `json:"username"`
	Password           string         `json:"password"`
	AuthType           enums.AuthType `json:"auth_type"`
	PrivateKey         string         `json:"private_key"`
	KeyPassword        string         `json:"key_password"`
	Description        string         `json:"description"`
	IsCommonCredential bool           `json:"is_common_credential"`
}

func (c *Credential) TableName() string {
	return "credentials"
}
