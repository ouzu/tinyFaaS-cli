package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/urfave/cli/v2"
)

func deleteCommand(c *cli.Context) error {

	res, err := http.Post(GetManagerUrl(c)+"/delete", "text/plain", strings.NewReader(c.Args().First()))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	s := string(b)

	if s == "" {
		fmt.Println("Ok")
	} else {
		return errors.New(s)
	}

	return nil
}
