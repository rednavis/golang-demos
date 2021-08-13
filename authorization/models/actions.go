package models

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Action struct {
	gorm.Model
	Action   string `json:"action"`
	DateTime string `json:"date_time"`
	UserId   uint   `json:"user_id"`
}

func (action *Action) Validate() (map[string]interface{}, bool) {

	if action.Action == "" {
		return u.Message(false, "User status should be on the payload"), false
	}

	if action.DateTime == "" {
		return u.Message(false, "DateTime should be on the payload"), false
	}

	if action.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (action *Action) Create() (map[string]interface{}) {

	if resp, ok := action.Validate(); !ok {
		return resp
	}

	GetDB().Create(action)

	resp := u.Message(true, "success")
	resp["action"] = action
	return resp
}

func GetActions(user uint) ([]*Action) {

	actions := make([]*Action, 0)
	err := GetDB().Table("actions").Where("user_id = ?", user).Find(&actions).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return actions
}
