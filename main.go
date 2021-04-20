package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

var config *AppConfigure

const curVersion = "0.2.1"

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
		Version: curVersion,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func server() {
	http.HandleFunc("/proxy", serverHandle)
	addr := fmt.Sprintf("%s:%d", config.Bind, config.Port)
	log.Println("http server start on ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

//http://myhpb.cn/proxy?content=base64(URL)
func serverHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.NotFound(w, req)
		return
	}
	log.Println("handle request: ", req.URL.String())
	content := req.URL.Query().Get("content")
	if content == "" {
		http.NotFound(w, req)
		return
	}
	b, err := fetchContent(content)
	if err != nil {
		log.Println("fetchContent err = ", err)
		http.Error(w, "fetch content error", 500)
		return
	}
	w.Header().Add("content-type", http.DetectContentType(b))
	w.Write(b)
}

func fetchContent(content string) ([]byte, error) {
	log.Println("fetchContent content: ", content)
	decoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Println("DecodeString err = ", err)
		return nil, err
	}
	contentURL := string(decoded)
	log.Println("content URL = ", contentURL)
	contentLocalPath := config.SavePath + "/" + content
	if exists(contentLocalPath) { //fetch in local dir
		log.Println("fetch from local, content: ", content)
		b, err := ioutil.ReadFile(contentLocalPath)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
	log.Println("fetch from network, content: ", content)
	b, err := httpGetHelper(contentURL)
	if err != nil {
		return nil, err
	}
	// save to local
	if err = writeToFile(contentLocalPath, b); err != nil {
		log.Println("write file error, err = ", err)
	}
	log.Println("fetchContent finish")
	return b, nil
}
