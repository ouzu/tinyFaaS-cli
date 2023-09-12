package main

import (
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func listCommand(c *cli.Context) error {
	res, err := http.Get(GetManagerUrl(c) + "/list")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
	return nil
}
