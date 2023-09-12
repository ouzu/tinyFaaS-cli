package main

import (
	"errors"
	"fmt"

	"github.com/pfandzelter/go-coap"
	"github.com/urfave/cli/v2"
)

func runCommand(c *cli.Context) error {
	command := c.Args().First()

	conn, err := coap.Dial("udp", GetCoapUrl(c))
	if err != nil {
		return err
	}

	req := coap.Message{
		Type: coap.Confirmable,
		Code: coap.GET,
	}

	req.SetPathString(command)

	res, err := conn.Send(req)
	if err != nil {
		return err
	}

	if res.Code != coap.Content {
		return errors.New(fmt.Sprintf("coap status %s", res.Code))
	}

	fmt.Println(string(res.Payload))

	return nil
}
