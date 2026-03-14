package cmd

import (
	"encoding/json"
	"fmt"
	"notion_go/base"
	"os"
)

func Query() {
	ApiKey := os.Getenv("NOTION_API_KEY")
	DatabaseID := os.Getenv("DATA_SOURCE_ID")

	prettyJSON := base.GetPages(DatabaseID, ApiKey)

	formatted, _ := json.MarshalIndent(prettyJSON, "", "  ")
	fmt.Println(string(formatted))
}
