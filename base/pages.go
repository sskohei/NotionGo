package base

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetPages(DatabaseID string, ApiKey string) []map[string]interface{} {

	url := "https://api.notion.com/v1/data_sources/" + DatabaseID + "/query"

	reqBody := []byte(`{}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+ApiKey)
	req.Header.Set("Notion-Version", "2025-09-03")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
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

	return pages
}
