// @author:zfy
// @date:2023/8/20 19:21

package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "goBlog"
	indexData.Description = "入门"
	jsonStr, _ := json.Marshal(indexData)

	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "goBlog"
	indexData.Description = "入门"

	t := template.New("index.html")
	path, _ := os.Getwd()
	//fmt.Println(path)
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8085",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
