package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/urfave/cli/v2"
)

func logsCommand(c *cli.Context) error {
	res, err := http.Get(GetManagerUrl(c) + "/logs")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Print(string(b))

	return nil
}
