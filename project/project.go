package project

import (
	"encoding/json"
	"fmt"
	"gitlab/api"
)

// Project 项目信息
type Project struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PathWithNamespace string `json:"path_with_namespace"`
	WebURL            string `json:"web_url"`
	SSHURL            string `json:"ssh_url_to_repo"`
	HTTPURL           string `json:"http_url_to_repo"`
}

// Merge merge request信息
type Merge struct {
	ID int `json:"id"`
}

func init() {

}

// Search 获取项目Id
func Search(name string) ([]*Project, error) {
	var data []interface{}
	params := make(map[string]interface{})
	params["search"] = name
	params["scope"] = "projects"
	err := api.GET("/api/v4/search", params, &data)
	if err != nil {
		return nil, err
	}

	var list []*Project
	if len(data) == 0 {
		return nil, nil
	}

	for i := 0; i < len(data); i++ {
		value, ok := data[i].(map[string]interface{})
		if ok {
			_data, _ := json.Marshal(value)
			var item Project
			err = json.Unmarshal([]byte(_data), &item)
			if err == nil {
				fmt.Printf("%d %s %s %s\n", item.ID, item.Name, item.PathWithNamespace, item.WebURL)
				list = append(list, &item)
			} else {
				return nil, err
			}
		}
	}
	return list, err
}

func SearchMergeByProjectName(project string, search string) ([]*Merge, error) {

	return nil, nil
}

func SearchMergeByProjectID(project string, search string) ([]*Merge, error) {

	return nil, nil
}

// SearchMergeRequest 搜索指定项目的merge request
func SearchMergeRequest(projectID string, search string) ([]*Project, error) {
	var data []interface{}
	params := make(map[string]interface{})
	params["search"] = search
	params["scope"] = "merge_requests"
	path := fmt.Sprintf("/api/v4/projects/%s/search", projectID)
	err := api.GET(path, params, &data)
	if err != nil {
		return nil, err
	}

	fmt.Printf("format result: %v\n", data)

	var list []*Project
	// if len(data) == 0 {
	// 	return nil, nil
	// }

	// for i := 0; i < len(data); i++ {
	// 	value, ok := data[i].(map[string]interface{})
	// 	if ok {
	// 		_data, _ := json.Marshal(value)
	// 		var item Project
	// 		err = json.Unmarshal([]byte(_data), &item)
	// 		if err == nil {
	// 			fmt.Printf("%d %s %s %s\n", item.ID, item.Name, item.PathWithNamespace, item.WebURL)
	// 			list = append(list, &item)
	// 		} else {
	// 			return nil, err
	// 		}
	// 	}
	// }
	return list, err
}
