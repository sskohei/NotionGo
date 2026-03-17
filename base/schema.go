package base

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"notion_go/model"
	"sort"
)

func GetSchema(DatabaseID string, ApiKey string) []model.Column {

	if DatabaseID == "" {
		fmt.Println("Datasource_idが指定されていません")
	}
	if ApiKey == "" {
		fmt.Println("ApiKeyが指定されていません")
	}

	url := "https://api.notion.com/v1/data_sources/" + DatabaseID

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+ApiKey)
	req.Header.Set("Notion-Version", "2025-09-03")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	props := result["properties"].(map[string]interface{})

	var columns []model.Column

	for name, v := range props {
		prop := v.(map[string]interface{})
		colType := prop["type"].(string)
		if colType == "select" || colType == "multi_select" || colType == "status" {
			selectInfo := prop[colType].(map[string]interface{})
			options := selectInfo["options"].([]interface{})

			var opts []string

			for _, opt := range options {
				o := opt.(map[string]interface{})
				opts = append(opts, o["name"].(string))
			}

			columns = append(columns, model.Column{
				Name:    name,
				Type:    colType,
				Options: opts,
			})
			continue
		}
		if colType == "title" {
			original := columns
			columns = append([]model.Column{{
				Name: name,
				Type: colType,
			}}, original...)
			continue
		} else {
			columns = append(columns, model.Column{
				Name: name,
				Type: colType,
			})
			continue
		}
	}

	middle := columns[1:]

	sort.Slice(middle, func(i, j int) bool {
		return middle[i].Name < middle[j].Name
	})

	original := middle
	middle = append([]model.Column{{
		Name:    columns[0].Name,
		Type:    columns[0].Type,
		Options: columns[0].Options,
	}}, original...)
	columns = middle

	return columns
}
