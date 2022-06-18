package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

type listResponse struct {
	Name     string `json:"name"`
	Hash     string `json:"hash"`
	Threads  int    `json:"threads"`
	Resource string `json:"resource"`
}

func listCommand(c *cli.Context) error {
	res, err := http.Get(BASE_URL + "/list")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var list []listResponse

	err = json.NewDecoder(res.Body).Decode(&list)
	if err != nil {
		return err
	}

	for i, f := range list {
		if i > 0 {
			fmt.Println()
		}

		fmt.Println("Name:    ", f.Name)
		fmt.Println("Hash:    ", f.Hash)
		fmt.Println("Threads: ", f.Threads)
		fmt.Println("Resource:", f.Resource)
	}

	return nil
}
