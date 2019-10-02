package main

import (
	"code.gitea.io/sdk/gitea"
	"encoding/json"
	"github.com/gobuffalo/packr"
	"github.com/ricardolonga/jsongo"
	"io/ioutil"
	"net/http"
)

type Config struct {
	AppName             string `json:"appName"`
	Subtitle            string `json:"subtitle"`
	APIAddress          string `json:"api_address"`
	GiteaServer         string `json:"gitea_server"`
	Token               string `json:"token"`
	User                string `json:"user"`
	Repo                string `json:"repo"`
	NotificationTitle   string `json:"notification_title"`
	NotificationMessage string `json:"notification_message"`
	Port                string `json:"port"`
}

func main() {
	box := packr.NewBox("./build")

	content, e := ioutil.ReadFile("./config.json")
	if e != nil {
		panic(e)
		return
	}
	c := &Config{}
	e = json.Unmarshal(content, c)
	if e != nil {
		panic(e)
		return
	}

	http.HandleFunc("/api/get-files", func(writer http.ResponseWriter, request *http.Request) {
		client := gitea.NewClient(c.GiteaServer, c.Token)
		r, i := client.ListReleases(c.User, c.Repo)
		base := jsongo.Array()
		if i != nil {
			panic(i)
			return
		}

		for _, v := range r {
			array := jsongo.Array()
			for _, attachment := range v.Attachments {
				array.Put(jsongo.Object().Put("name", attachment.Name).Put("downloadURL", attachment.DownloadURL))
			}
			title := v.Title
			if len(title) == 0 {
				title = "无标题"
			}
			base.Put(jsongo.Object().Put("title", title).Put("files", array))
		}
		_, i = writer.Write([]byte(base.String()))
		if i != nil {
			panic(i)
		}
	})
	http.Handle("/", http.FileServer(box))
	e = http.ListenAndServe(":"+c.Port, nil)
	if e != nil {
		println("Port unavailable")
	}
}
