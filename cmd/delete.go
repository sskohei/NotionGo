package cmd

import (
	"flag"
	"notion_go/base"
	"os"
)

func Delete(text []string) {
	ApiKey := os.Getenv("NOTION_API_KEY")
	DatabaseID := os.Getenv("DATA_SOURCE_ID")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	keyword := deleteCmd.String("k", "", "キーワード")
	property := deleteCmd.String("p", "", "プロパティ")
	deleteCmd.Parse(text)

	schema := base.GetSchema(DatabaseID, ApiKey)

	base.DeleteData(DatabaseID, ApiKey, schema, *keyword, *property)

	pages := base.GetPages(DatabaseID, ApiKey)

	base.RenderTable(schema, pages)
}
