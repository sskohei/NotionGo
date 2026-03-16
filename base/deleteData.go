package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"notion_go/model"
)

func DeleteData(DatabaseID string, ApiKey string, Schema []model.Column, Keyword string, Property string) {
	pageData := Filter("equals", DatabaseID, ApiKey, Schema, Keyword, Property)
	pageID := pageData[0]["id"].(string)

	condition := map[string]interface{}{
		"parent": map[string]interface{}{
			"data_source_id": DatabaseID,
		},
		"archived": true,
	}

	conditionBytes, _ := json.Marshal(condition)

	url := "https://api.notion.com/v1/pages/" + pageID

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(conditionBytes))
	req.Header.Set("Authorization", "Bearer "+ApiKey)
	req.Header.Set("Notion-Version", "2025-09-03")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Response body:")
	fmt.Println(string(body))

}
