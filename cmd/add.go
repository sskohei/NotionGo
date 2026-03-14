package cmd

import (
	"notion_go/base"
	"os"
)

func Add() {
	ApiKey := os.Getenv("NOTION_API_KEY")
	DatabaseID := os.Getenv("DATA_SOURCE_ID")
	schema := base.GetSchema(DatabaseID, ApiKey)
	base.Add(DatabaseID, ApiKey, schema)

	pages := base.GetPages(DatabaseID, ApiKey)

	base.RenderTable(schema, pages)
}
