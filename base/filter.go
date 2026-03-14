package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"notion_go/model"
)

func Filter(option string, DatabaseID string, ApiKey string, Schema []model.Column, Keyword string, Property string) []map[string]interface{} {
	var t string
	for _, col := range Schema {
		if col.Name == Property {
			t = col.Type
			break
		}
	}

	if t == "" {
		return nil
	}

	condition := map[string]interface{}{
		"filter": map[string]interface{}{
			"property": Property,
			t: map[string]interface{}{
				option: Keyword,
			},
		},
	}

	conditionBytes, _ := json.Marshal(condition)

	url := "https://api.notion.com/v1/data_sources/" + DatabaseID + "/query"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(conditionBytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+ApiKey)
	req.Header.Set("Notion-Version", "2025-09-03")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	rawPages := result["results"].([]interface{})

	var pages []map[string]interface{}

	for _, p := range rawPages {
		pages = append(pages, p.(map[string]interface{}))
	}

	fmt.Println(condition)

	return pages
}
