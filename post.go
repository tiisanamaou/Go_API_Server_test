package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostResponse struct {
	UserID   string `json:"UseID"`
	UserRank int    `json:"UserRank"`
	UserName string `json:"UserName"`
}

func post(w http.ResponseWriter, r *http.Request) {
	//Validate request
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add Response
	postResponse := PostResponse{
		UserID:   "0001",
		UserRank: 10,
		UserName: "maomao",
	}
	fmt.Println(postResponse)
	jsonData, err := json.Marshal(postResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonData)
	fmt.Println(string(jsonData))

	len := r.ContentLength
	body := make([]byte, len) // Content-Length と同じサイズの byte 配列を用意
	r.Body.Read(body)         // byte 配列にリクエストボディを読み込む
	//fmt.Fprintln(w, string(body))
	fmt.Fprintln(w, string(jsonData))
}
