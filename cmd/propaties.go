package cmd

import (
	"fmt"
	"notion_go/base"
	"os"
)

func Propaties() {
	ApiKey := os.Getenv("NOTION_API_KEY")
	DatabaseID := os.Getenv("DATA_SOURCE_ID")

	schema := base.GetSchema(DatabaseID, ApiKey)
	fmt.Println("=== Database Schema ===")
	for _, col := range schema {
		fmt.Printf("Column: %-20s Type: %s\n Options: %s\n", col.Name, col.Type, col.Options)
	}
}
