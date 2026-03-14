package cmd

import (
	"flag"
	"fmt"
	"notion_go/base"
	"os"
)

func Equal(text []string) {
	ApiKey := os.Getenv("NOTION_API_KEY")
	DatabaseID := os.Getenv("DATA_SOURCE_ID")

	filterCmd := flag.NewFlagSet("filter", flag.ExitOnError)
	keyword := filterCmd.String("k", "", "キーワード")
	property := filterCmd.String("p", "", "プロパティ")
	filterCmd.Parse(text)

	if *keyword == "" {
		fmt.Println("キーワードを-kで指定してください")
		return
	}

	if *property == "" {
		fmt.Println("-pでプロパティを指定してください")
		return
	}

	// ① スキーマ取得
	schema := base.GetSchema(DatabaseID, ApiKey)

	// ② データ取得
	pages := base.Filter("equals", DatabaseID, ApiKey, schema, *keyword, *property)

	if pages == nil {
		fmt.Println("検索したものと一致するものはありませんでした")
		return
	}

	// ③ 表示
	base.RenderTable(schema, pages)
}
