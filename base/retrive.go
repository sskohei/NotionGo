package base

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Retrive(DatabaseID string, ApiKey string) {
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

	pretty, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pretty))
}
