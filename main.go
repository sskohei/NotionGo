package main

import (
	"fmt"
	"notion_go/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使い方")
		fmt.Println(" notion query")
		return
	}

	switch os.Args[1] {

	case "query":
		cmd.Query()

	case "equal":
		cmd.Equal(os.Args[2:])

	case "contain":
		cmd.Contain(os.Args[2:])

	case "properties":
		cmd.Properties()

	case "list":
		cmd.List()

	case "add":
		cmd.Add()

	case "delete":
		cmd.Delete(os.Args[2:])

	case "test":
		cmd.Test()

	default:
		return
	}
}
