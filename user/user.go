package user

import (
	"encoding/json"
	"fmt"
	"gitlab/api"
)

// User 用户信息
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`     // Name 展示名，可修改
	UserName string `json:"username"` // 账号名，不可更改
	State    string `json:"state"`
	Avatar   string `json:"avatar_url"`
	WebURL   string `json:"web_url"`
}

// Search 搜索用户
func Search(name string) ([]*User, error) {
	var data []interface{}
	params := make(map[string]interface{})
	params["search"] = name
	params["scope"] = "users"
	err := api.GET("/api/v4/search", params, &data)
	if err != nil {
		return nil, err
	}

	var list []*User
	if len(data) == 0 {
		return nil, nil
	}

	for i := 0; i < len(data); i++ {
		value, ok := data[i].(map[string]interface{})
		if ok {
			_data, _ := json.Marshal(value)
			var item User
			err = json.Unmarshal([]byte(_data), &item)
			if err == nil {
				fmt.Printf("%d %s %s %s\n", item.ID, item.Name, item.UserName, item.WebURL)
				list = append(list, &item)
			} else {
				return nil, err
			}
		}
	}
	return list, err
}
