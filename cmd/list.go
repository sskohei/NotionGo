package cmd

import (
	"notion_go/base"
	"os"
)

func List() {

	ApiKey := os.Getenv("NOTION_API_KEY")
	DatabaseID := os.Getenv("DATA_SOURCE_ID")

	// ① スキーマ取得
	schema := base.GetSchema(DatabaseID, ApiKey)

	// ② データ取得
	pages := base.GetPages(DatabaseID, ApiKey)

	// ③ 表示
	base.RenderTable(schema, pages)
}
