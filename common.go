package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

const timeOutDur = 10

func replaceHost(str string, host string) string {
	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}
	return host + u.Path
}

func httpGetHelper(str string) ([]byte, error) {
	timeout := time.Duration(timeOutDur * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", str, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "f0568dd4-7afb-8703-33ec-83be5559a95d")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func exists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		return os.IsExist(err)
	}
	return true
}

func writeToFile(fileName string, data []byte) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt(data, n)
		defer f.Close()
	}
	return err
}
