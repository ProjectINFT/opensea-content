package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/urfave/cli/v2"
)

var config *AppConfigure

func main() {
	configPath := ""
	app := &cli.App{
		Usage: "opensea asset image proxy",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Value:       defaultConfigPath,
				Usage:       "set app configure file",
				Destination: &configPath,
			},
		},
		Action: func(c *cli.Context) error {
			config, _ = loadConfigure(configPath)
			server()
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func server() {
	http.HandleFunc("/", serverHandle)
	addr := fmt.Sprintf("%s:%d", config.Bind, config.Port)
	log.Println("http server start on ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func serverHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.NotFound(w, req)
		return
	}
	str := req.URL.String()
	log.Println("handle request: ", str)
	b, err := fetchContent(str)
	if err != nil {
		http.Error(w, "fetch content error", 500)
		return
	}
	w.Write(b)
}

func fetchContent(content string) ([]byte, error) {
	u, err := url.Parse(content)
	if err != nil {
		return nil, err
	}
	path := config.SavePath + u.Path
	if exists(path) { //fetch in local dir
		log.Println("fetch from local, content: ", u.Path)
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
	log.Println("fetch from network, content: ", u.Path)
	b, err := httpGetHelper(content)
	if err != nil {
		return nil, err
	}
	// save to local
	if err = writeToFile(path, b); err != nil {
		log.Println("write file error, err = ", err)
	}
	return b, nil
}
