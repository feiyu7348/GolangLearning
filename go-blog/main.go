// @author:zfy
// @date:2023/8/20 19:21

package main

import (
	"encoding/json"
	"log"
	"net/http"
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

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8085",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
