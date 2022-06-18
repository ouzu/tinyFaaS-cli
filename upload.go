package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

type uploadRequest struct {
	Name    string `json:"name"`
	Threads int    `json:"threads"`
	Zip     string `json:"zip"`
}

func uploadCommand(c *cli.Context) error {
	if c.Args().Len() != 3 {
		return fmt.Errorf("Need exactly 3 arguments.")
	}

	dir := c.Args().Get(0)
	name := c.Args().Get(1)
	threads, err := strconv.Atoi(c.Args().Get(2))

	var zipBuffer bytes.Buffer
	zipWriter := zip.NewWriter(bufio.NewWriter(&zipBuffer))

	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			stripped := strings.TrimPrefix(path, dir)

			if !info.IsDir() {
				fmt.Fprintln(os.Stderr, "adding", stripped)

				fileWriter, err := zipWriter.Create(stripped)
				if err != nil {
					return err
				}

				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				_, err = io.Copy(fileWriter, file)
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}

	zipWriter.Close()

	req := uploadRequest{
		Name:    name,
		Threads: threads,
		Zip:     base64.StdEncoding.EncodeToString(zipBuffer.Bytes()),
	}

	j, err := json.Marshal(req)
	if err != nil {
		return err
	}

	res, err := http.Post(BASE_URL+"/upload", "text/plain", bytes.NewReader(j))
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
