package base

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"notion_go/model"
	"os"
	"strconv"
	"strings"
)

func Add(DatabaseID string, ApiKey string, Schema []model.Column) {
	properties := map[string]interface{}{}

	for _, col := range Schema {
		if col.Type == "date" || col.Type == "formula" {
			continue
		}

		if col.Type == "select" || col.Type == "multi_select" {
			input := selectPrompt(col)
			if input == "" {
				continue
			}
			properties[col.Name] = designate(col, input)
		} else {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println(col.Name + ": ")
			input, _ := reader.ReadString('\n')

			input = strings.TrimSpace(input)

			if input == "" {
				continue
			}

			properties[col.Name] = designate(col, input)
		}
	}

	condition := map[string]interface{}{
		"parent": map[string]interface{}{
			"data_source_id": DatabaseID,
		},
		"properties": properties,
	}

	conditionBytes, _ := json.Marshal(condition)

	url := "https://api.notion.com/v1/pages"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(conditionBytes))
	req.Header.Set("Authorization", "Bearer "+ApiKey)
	req.Header.Set("Notion-Version", "2025-09-03")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Response body:")
	fmt.Println(string(body))

	fmt.Println(properties)
}

func designate(m model.Column, word string) map[string]interface{} {
	switch m.Type {
	case "title":
		return map[string]interface{}{
			"title": []map[string]interface{}{
				{
					"text": map[string]interface{}{
						"content": word,
					},
				},
			},
		}

	case "people":
		return map[string]interface{}{
			"people": []map[string]interface{}{
				{
					"id": word,
				},
			},
		}

	case "multi_select":
		return map[string]interface{}{
			"multi_select": []map[string]interface{}{
				{
					"name": word,
				},
			},
		}

	case "number":
		num, err := strconv.Atoi(word)
		if err != nil {
			fmt.Println("数字ではありません", err)
			return nil
		}
		return map[string]interface{}{
			m.Type: num,
		}

	case "status":
		return map[string]interface{}{
			"status": map[string]interface{}{
				"name": word,
			},
		}

	case "select":
		return map[string]interface{}{
			"select": map[string]interface{}{
				"name": word,
			},
		}

	default:
		return nil
	}
}

func selectPrompt(sel model.Column) string {
	fmt.Println(sel.Name + "を選択してください")
	for i, opt := range sel.Options {
		fmt.Printf("%d: %s\n", i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(sel.Options) {
			fmt.Println("番号を入力してください")
			continue
		}

		return sel.Options[index-1]
	}
}
