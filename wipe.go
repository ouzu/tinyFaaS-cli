package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/urfave/cli/v2"
)

func wipeCommand(c *cli.Context) error {
	res, err := http.Post(GetManagerUrl(c)+"/wipe", "text/plain", bytes.NewReader([]byte{}))
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
