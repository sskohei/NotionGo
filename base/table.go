package base

import (
	"fmt"
	"notion_go/model"
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderTable(schema []model.Column, pages []map[string]interface{}) {

	table := tablewriter.NewWriter(os.Stdout)

	// ヘッダー作成
	var headers []string
	for _, col := range schema {
		headers = append(headers, col.Name)
	}
	table.Header(headers)

	// 行データ作成
	for _, page := range pages {

		props := page["properties"].(map[string]interface{})

		var row []string

		for _, col := range schema {
			prop := props[col.Name].(map[string]interface{})
			value := extractValue(prop)
			row = append(row, value)
		}

		table.Append(row)
	}

	table.Render()
}

func extractValue(prop map[string]interface{}) string {

	t := prop["type"].(string)

	switch t {

	case "title":
		return joinText(prop["title"], "plain_text")

	case "rich_text":
		return joinText(prop["rich_text"], "plain_text")

	case "people":
		return joinText(prop["people"], "name")

	case "multi_select":
		return joinText(prop["multi_select"], "name")

	case "number":
		if prop["number"] == nil {
			return ""
		}
		return fmt.Sprintf("%v", prop["number"])

	case "status":
		if prop["status"] == nil {
			return ""
		}
		return prop["status"].(map[string]interface{})["name"].(string)

	case "select":
		if prop["select"] == nil {
			return ""
		}
		return prop["select"].(map[string]interface{})["name"].(string)

	case "formula":
		if prop["formula"] == nil {
			return ""
		}
		return fmt.Sprintf("%v", prop["formula"].(map[string]interface{})["number"].(float64))

	case "date":
		if prop["date"] == nil {
			return ""
		}
		return prop["date"].(map[string]interface{})["start"].(string)

	default:
		return ""
	}
}

func joinText(v interface{}, t string) string {

	arr, ok := v.([]interface{})
	if !ok {
		return ""
	}

	result := ""

	for _, item := range arr {
		m := item.(map[string]interface{})
		result += m[t].(string)
	}

	return result
}
