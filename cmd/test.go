package cmd

import (
	"notion_go/base"
	"os"
)

func Test() {
	ApiKey := os.Getenv("NOTION_API_KEY")
	DatabaseID := os.Getenv("DATA_SOURCE_ID")

	base.Retrive(DatabaseID, ApiKey)
}
